/*
Copyright(c) 2023-present Accton. All rights reserved. www.accton.com.tw
*/

package diskvolume

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	storagev1 "k8s.io/api/storage/v1"

	virtzv1alpha1 "kubesphere.io/api/virtualization/v1alpha1"
)

const (
	controllerName        = "diskvolume-controller"
	successSynced         = "Synced"
	messageResourceSynced = "DiskVolume synced successfully"
	pvcNamePrefix         = "tpl-"
)

// Reconciler reconciles a disk volume object
type Reconciler struct {
	client.Client
	Logger                  logr.Logger
	Recorder                record.EventRecorder
	MaxConcurrentReconciles int
}

func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	if r.Client == nil {
		r.Client = mgr.GetClient()
	}
	if r.Logger == nil {
		r.Logger = ctrl.Log.WithName("controllers").WithName(controllerName)
	}
	if r.Recorder == nil {
		r.Recorder = mgr.GetEventRecorderFor(controllerName)
	}
	if r.MaxConcurrentReconciles <= 0 {
		r.MaxConcurrentReconciles = 1
	}
	return ctrl.NewControllerManagedBy(mgr).
		Named(controllerName).
		WithOptions(controller.Options{
			MaxConcurrentReconciles: r.MaxConcurrentReconciles,
		}).
		For(&virtzv1alpha1.DiskVolume{}).
		Complete(r)
}

func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	klog.V(2).Infof("Reconciling VirtualMachine %s/%s", req.Namespace, req.Name)

	rootCtx := context.Background()
	dv := &virtzv1alpha1.DiskVolume{}
	if err := r.Get(rootCtx, req.NamespacedName, dv); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// get default storage class name
	scName := ""
	scList := &storagev1.StorageClassList{}
	if err := r.List(rootCtx, scList); err != nil {
		return ctrl.Result{}, err
	}
	for _, sc := range scList.Items {
		if sc.Annotations["storageclass.kubernetes.io/is-default-class"] == "true" {
			scName = sc.Name
			break
		}
	}
	if scName == "" {
		return ctrl.Result{}, fmt.Errorf("no default storage class found")
	}

	dv_instance := dv.DeepCopy()

	status := &dv_instance.Status
	if !status.Created {
		// create pvc
		if dv_instance.Spec.Source.Blank != nil {
			err := r.createPVC(dv_instance, scName)
			if err != nil {
				return ctrl.Result{}, err
			}
		}

		status.Created = true

	}

	if !reflect.DeepEqual(dv, dv_instance) {
		if err := r.Update(rootCtx, dv_instance); err != nil {
			return ctrl.Result{}, err
		}
	}

	// update status
	status.Ready = true
	if err := r.Status().Update(rootCtx, dv_instance); err != nil {
		return ctrl.Result{}, err
	}

	// update event
	r.Recorder.Event(dv, corev1.EventTypeNormal, successSynced, messageResourceSynced)
	return ctrl.Result{}, nil

}

func (r *Reconciler) createPVC(dv_instance *virtzv1alpha1.DiskVolume, scName string) error {
	klog.Infof("Creating pvc %s/%s", dv_instance.Namespace, dv_instance.Spec.PVCName)

	blockOwnerDeletion := true
	controller := true

	pvc := &corev1.PersistentVolumeClaim{}
	pvc.Name = pvcNamePrefix + dv_instance.Name
	pvc.Namespace = dv_instance.Namespace
	pvc.Spec.AccessModes = []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce}
	pvc.Spec.Resources = corev1.ResourceRequirements{}
	pvc.Spec.Resources.Requests = corev1.ResourceList{}
	pvc.Spec.Resources.Requests[corev1.ResourceStorage] = dv_instance.Spec.Resources.Requests[corev1.ResourceStorage]
	pvc.Spec.StorageClassName = &scName
	// owner reference
	pvc.OwnerReferences = []metav1.OwnerReference{
		{
			APIVersion:         dv_instance.APIVersion,
			BlockOwnerDeletion: &blockOwnerDeletion,
			Controller:         &controller,
			Kind:               dv_instance.Kind,
			Name:               dv_instance.Name,
			UID:                dv_instance.UID,
		},
	}

	if err := r.Create(context.Background(), pvc); err != nil {
		return err
	}

	klog.Infof("PVC %s/%s created", pvc.Namespace, pvc.Name)

	return nil
}
