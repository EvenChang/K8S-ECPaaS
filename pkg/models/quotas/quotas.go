/*
Copyright 2019 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package quotas

import (
	"context"

	"github.com/minio/minio-go/v7"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"k8s.io/klog"
	"kubesphere.io/kubesphere/pkg/api"

	ksinformers "kubesphere.io/kubesphere/pkg/client/informers/externalversions"
)

const (
	podsKey                   = "count/pods"
	daemonsetsKey             = "count/daemonsets.apps"
	deploymentsKey            = "count/deployments.apps"
	ingressKey                = "count/ingresses.extensions"
	servicesKey               = "count/services"
	statefulsetsKey           = "count/statefulsets.apps"
	persistentvolumeclaimsKey = "persistentvolumeclaims"
	jobsKey                   = "count/jobs.batch"
	cronJobsKey               = "count/cronjobs.batch"
	s2iBuilders               = "count/s2ibuilders.devops.kubesphere.io"
	virtualmachinesKey        = "count/virtualmachines.virtualization.ecpaas.io"
	imagetemplatesKey         = "count/images.virtualization.ecpaas.io"
	diskvolumesKey            = "count/disks.virtualization.ecpaas.io"
	minioImagesKey            = "count/files.virtualization.ecpaas.io"
)

var supportedResources = map[string]schema.GroupVersionResource{
	deploymentsKey:            {Group: "apps", Version: "v1", Resource: "deployments"},
	daemonsetsKey:             {Group: "apps", Version: "v1", Resource: "daemonsets"},
	statefulsetsKey:           {Group: "apps", Version: "v1", Resource: "statefulsets"},
	podsKey:                   {Group: "", Version: "v1", Resource: "pods"},
	servicesKey:               {Group: "", Version: "v1", Resource: "services"},
	persistentvolumeclaimsKey: {Group: "", Version: "v1", Resource: "persistentvolumeclaims"},
	ingressKey:                {Group: "networking.k8s.io", Version: "v1", Resource: "ingresses"},
	jobsKey:                   {Group: "batch", Version: "v1", Resource: "jobs"},
	cronJobsKey:               {Group: "batch", Version: "v1beta1", Resource: "cronjobs"},
	s2iBuilders:               {Group: "devops.kubesphere.io", Version: "v1alpha1", Resource: "s2ibuilders"},
}

var supportedVirtualizationResources = map[string]schema.GroupVersionResource{
	virtualmachinesKey: {Group: "virtualization.ecpaas.io", Version: "v1alpha1", Resource: "virtualmachines"},
	imagetemplatesKey:  {Group: "virtualization.ecpaas.io", Version: "v1alpha1", Resource: "imagetemplates"},
	diskvolumesKey:     {Group: "virtualization.ecpaas.io", Version: "v1alpha1", Resource: "diskvolumes"},
}

type ResourceQuotaGetter interface {
	GetClusterQuota() (*api.ResourceQuota, error)
	GetNamespaceQuota(namespace string) (*api.NamespacedResourceQuota, error)
	GetVirtualizationNamespaceQuota(namespace string, minioClient *minio.Client) (*api.NamespacedResourceQuota, error)
}

type resourceQuotaGetter struct {
	k8sInformers        informers.SharedInformerFactory
	kubesphereInformers ksinformers.SharedInformerFactory
}

func NewResourceQuotaGetter(k8sInformers informers.SharedInformerFactory, kubesphereInformers ksinformers.SharedInformerFactory) ResourceQuotaGetter {
	return &resourceQuotaGetter{
		k8sInformers:        k8sInformers,
		kubesphereInformers: kubesphereInformers,
	}
}

func (c *resourceQuotaGetter) getUsage(namespace, resource string) (int, error) {

	genericInformer, err := c.k8sInformers.ForResource(supportedResources[resource])
	if err != nil {
		// we deliberately ignore error if trying to get non existed resource
		return 0, nil
	}

	result, err := genericInformer.Lister().ByNamespace(namespace).List(labels.Everything())
	if err != nil {
		return 0, err
	}

	return len(result), nil
}

func (c *resourceQuotaGetter) getVirtualizationUsage(namespace, resource string) (int, error) {

	genericInformer, err := c.kubesphereInformers.ForResource(supportedVirtualizationResources[resource])
	if err != nil {
		// we deliberately ignore error if trying to get non existed resource
		return 0, nil
	}

	result, err := genericInformer.Lister().ByNamespace(namespace).List(labels.Everything())
	if err != nil {
		return 0, err
	}

	return len(result), nil
}

// no one use this api anymore， marked as deprecated
func (c *resourceQuotaGetter) GetClusterQuota() (*api.ResourceQuota, error) {

	quota := v1.ResourceQuotaStatus{Hard: make(v1.ResourceList), Used: make(v1.ResourceList)}

	for r := range supportedResources {
		used, err := c.getUsage("", r)
		if err != nil {
			return nil, err
		}
		var quantity resource.Quantity
		quantity.Set(int64(used))
		quota.Used[v1.ResourceName(r)] = quantity
	}

	return &api.ResourceQuota{Namespace: "\"\"", Data: quota}, nil

}

func (c *resourceQuotaGetter) GetNamespaceQuota(namespace string) (*api.NamespacedResourceQuota, error) {
	quota, err := c.getNamespaceResourceQuota(namespace)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	if quota == nil {
		quota = &v1.ResourceQuotaStatus{Hard: make(v1.ResourceList), Used: make(v1.ResourceList)}
	}

	var resourceQuotaLeft = v1.ResourceList{}

	for key, hardLimit := range quota.Hard {
		if used, ok := quota.Used[key]; ok {
			left := hardLimit.DeepCopy()
			left.Sub(used)
			if hardLimit.Cmp(used) < 0 {
				left = resource.MustParse("0")
			}

			resourceQuotaLeft[key] = left
		}
	}

	// add extra quota usage, cause user may not specify them
	for key := range supportedResources {
		// only add them when they don't exist in quotastatus
		if _, ok := quota.Used[v1.ResourceName(key)]; !ok {
			used, err := c.getUsage(namespace, key)
			if err != nil {
				klog.Error(err)
				return nil, err
			}

			quota.Used[v1.ResourceName(key)] = *(resource.NewQuantity(int64(used), resource.DecimalSI))
		}
	}

	var result = api.NamespacedResourceQuota{
		Namespace: namespace,
	}
	result.Data.Hard = quota.Hard
	result.Data.Used = quota.Used
	result.Data.Left = resourceQuotaLeft

	return &result, nil

}

func (c *resourceQuotaGetter) GetVirtualizationNamespaceQuota(namespace string, minioClient *minio.Client) (*api.NamespacedResourceQuota, error) {
	quota, err := c.getNamespaceResourceQuota(namespace)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	if quota == nil {
		quota = &v1.ResourceQuotaStatus{Hard: make(v1.ResourceList), Used: make(v1.ResourceList)}
	}

	var resourceQuotaLeft = v1.ResourceList{}

	for key, hardLimit := range quota.Hard {
		if used, ok := quota.Used[key]; ok {
			left := hardLimit.DeepCopy()
			left.Sub(used)
			if hardLimit.Cmp(used) < 0 {
				left = resource.MustParse("0")
			}

			resourceQuotaLeft[key] = left
		}
	}

	// add extra quota usage, cause user may not specify them
	for key := range supportedVirtualizationResources {
		// only add them when they don't exist in quotastatus
		if _, ok := quota.Used[v1.ResourceName(key)]; !ok {
			used, err := c.getVirtualizationUsage(namespace, key)
			if err != nil {
				klog.Error(err)
				return nil, err
			}

			quota.Used[v1.ResourceName(key)] = *(resource.NewQuantity(int64(used), resource.DecimalSI))
		}
	}

	minioImages := 0
	if minioClient != nil {
		bucketName := "ecpaas-images"
		minioServiceName := "minio"
		list := c.k8sInformers.Core().V1().Services().Lister()
		serviceList, err := list.Services("").List(labels.Everything())

		var minioService *v1.Service
		for _, service := range serviceList {
			if service.Name == minioServiceName {
				minioService = service
				break
			}
		}

		if minioService == nil {
			klog.Warning("Cannot find the minio service ", err)
		}

		objectCh := minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{})

		for object := range objectCh {
			if object.Err == nil {
				minioImages++
			}
		}

		quota.Used[minioImagesKey] = *(resource.NewQuantity(int64(minioImages), resource.DecimalSI))
	}

	var result = api.NamespacedResourceQuota{
		Namespace: namespace,
	}
	result.Data.Hard = quota.Hard
	result.Data.Used = quota.Used
	result.Data.Left = resourceQuotaLeft

	return &result, nil

}

func updateNamespaceQuota(tmpResourceList, resourceList v1.ResourceList) {
	if tmpResourceList == nil {
		tmpResourceList = resourceList
	}
	for res, usage := range resourceList {
		tmpUsage, exist := tmpResourceList[res]
		if !exist {
			tmpResourceList[res] = usage
		}
		if tmpUsage.Cmp(usage) == 1 {
			tmpResourceList[res] = usage
		}
	}
}

func (c *resourceQuotaGetter) getNamespaceResourceQuota(namespace string) (*v1.ResourceQuotaStatus, error) {
	resourceQuotaLister := c.k8sInformers.Core().V1().ResourceQuotas().Lister()
	quotaList, err := resourceQuotaLister.ResourceQuotas(namespace).List(labels.Everything())
	if err != nil {
		klog.Error(err)
		return nil, err
	} else if len(quotaList) == 0 {
		return nil, nil
	}

	quotaStatus := v1.ResourceQuotaStatus{Hard: make(v1.ResourceList), Used: make(v1.ResourceList)}

	for _, quota := range quotaList {
		updateNamespaceQuota(quotaStatus.Hard, quota.Status.Hard)
		updateNamespaceQuota(quotaStatus.Used, quota.Status.Used)
	}

	return &quotaStatus, nil
}
