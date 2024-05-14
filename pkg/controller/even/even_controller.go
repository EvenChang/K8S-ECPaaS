/*
Copyright(c) 2023-present Accton. All rights reserved. www.accton.com.tw
*/

package even

import (
	"context"
	"reflect"

	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubesphere.io/api/even/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const (
	controllerName        = "even-controller"
	successSynced         = "Synced"
	messageResourceSynced = "Even Resource synced successfully"
)

// Reconciler reconciles a image template object
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

	return ctrl.NewControllerManagedBy(mgr).
		Named(controllerName).
		For(&v1alpha1.Even{}).
		Complete(r)
}

func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	klog.V(2).Infof("Reconciling Even Resources %s/%s", req.Namespace, req.Name)

	rootCtx := context.Background()
	even := &v1alpha1.Even{}
	if err := r.Get(rootCtx, req.NamespacedName, even); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	even_instance := even.DeepCopy()

	deploymentname := even_instance.Spec.DeploymentName
	klog.Infof("Even's Deployment Name: %s\n", deploymentname)

	if NeedToAddFinalizer(even_instance, "finalizers.even.ecpaas.io") {
		klog.Infof("Adding finalizer for Even %s/%s", req.Namespace, req.Name)
		if err := r.addFinalizer(even_instance); err != nil {
			return ctrl.Result{}, err
		}
	}

	deployment := &appsv1.Deployment{}
	if err := r.Get(rootCtx, client.ObjectKey{Name: deploymentname, Namespace: "default"}, deployment); err != nil {

		klog.Infof("Even TEST")
		klog.Infof("Deployment not found %s", err)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	replicaValue := int32(even_instance.Spec.Replicas)
	deployment.Spec.Replicas = &replicaValue

	if !reflect.DeepEqual(deployment, even_instance) {
		if err := r.Update(rootCtx, deployment); err != nil {
			return ctrl.Result{}, err
		}
	}

	if IsDeletionCandidate(even_instance, "finalizers.even.ecpaas.io") {
		klog.Infof("Deleting Even %s/%s", req.Namespace, req.Name)

		if err := r.Delete(rootCtx, deployment); err != nil {
			return ctrl.Result{}, err
		}

		klog.Infof("Removing finalizer for Even %s/%s", req.Namespace, req.Name)
		if err := r.removeFinalizer(even_instance); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// IsDeletionCandidate checks if object is candidate to be deleted
func IsDeletionCandidate(obj metav1.Object, finalizer string) bool {
	return obj.GetDeletionTimestamp() != nil && ContainsString(obj.GetFinalizers(),
		finalizer, nil)
}

// NeedToAddFinalizer checks if need to add finalizer to object
func NeedToAddFinalizer(obj metav1.Object, finalizer string) bool {
	return obj.GetDeletionTimestamp() == nil && !ContainsString(obj.GetFinalizers(),
		finalizer, nil)
}

// ContainsString checks if a given slice of strings contains the provided string.
// If a modifier func is provided, it is called with the slice item before the comparation.
func ContainsString(slice []string, s string, modifier func(s string) string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
		if modifier != nil && modifier(item) == s {
			return true
		}
	}
	return false
}

func (c *Reconciler) addFinalizer(even *v1alpha1.Even) error {
	clone := even.DeepCopy()
	controllerutil.AddFinalizer(clone, "finalizers.even.ecpaas.io")

	err := c.Update(context.Background(), clone)
	if err != nil {
		klog.V(3).Infof("Error adding  finalizer to virtualmachine %s: %v", even.Name, err)
		return err
	}
	klog.V(3).Infof("Added finalizer to virtualmachine %s", even.Name)
	return nil
}

func (c *Reconciler) removeFinalizer(even *v1alpha1.Even) error {
	clone := even.DeepCopy()
	controllerutil.RemoveFinalizer(clone, "finalizers.even.ecpaas.io")
	err := c.Update(context.Background(), clone)
	if err != nil {
		klog.V(3).Infof("Error removing  finalizer from virtualmachine %s: %v", even.Name, err)
		return err
	}
	klog.V(3).Infof("Removed protection finalizer from virtualmachine %s", even.Name)
	return nil
}
