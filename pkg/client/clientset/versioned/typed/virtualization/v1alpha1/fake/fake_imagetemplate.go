/*
Copyright 2020 The KubeSphere Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubesphere.io/api/virtualization/v1alpha1"
)

// FakeImageTemplates implements ImageTemplateInterface
type FakeImageTemplates struct {
	Fake *FakeVirtualizationV1alpha1
	ns   string
}

var imagetemplatesResource = schema.GroupVersionResource{Group: "virtualization.ecpaas.io", Version: "v1alpha1", Resource: "imagetemplates"}

var imagetemplatesKind = schema.GroupVersionKind{Group: "virtualization.ecpaas.io", Version: "v1alpha1", Kind: "ImageTemplate"}

// Get takes name of the imageTemplate, and returns the corresponding imageTemplate object, and an error if there is any.
func (c *FakeImageTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ImageTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(imagetemplatesResource, c.ns, name), &v1alpha1.ImageTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageTemplate), err
}

// List takes label and field selectors, and returns the list of ImageTemplates that match those selectors.
func (c *FakeImageTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ImageTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(imagetemplatesResource, imagetemplatesKind, c.ns, opts), &v1alpha1.ImageTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ImageTemplateList{ListMeta: obj.(*v1alpha1.ImageTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ImageTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested imageTemplates.
func (c *FakeImageTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(imagetemplatesResource, c.ns, opts))

}

// Create takes the representation of a imageTemplate and creates it.  Returns the server's representation of the imageTemplate, and an error, if there is any.
func (c *FakeImageTemplates) Create(ctx context.Context, imageTemplate *v1alpha1.ImageTemplate, opts v1.CreateOptions) (result *v1alpha1.ImageTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(imagetemplatesResource, c.ns, imageTemplate), &v1alpha1.ImageTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageTemplate), err
}

// Update takes the representation of a imageTemplate and updates it. Returns the server's representation of the imageTemplate, and an error, if there is any.
func (c *FakeImageTemplates) Update(ctx context.Context, imageTemplate *v1alpha1.ImageTemplate, opts v1.UpdateOptions) (result *v1alpha1.ImageTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(imagetemplatesResource, c.ns, imageTemplate), &v1alpha1.ImageTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImageTemplates) UpdateStatus(ctx context.Context, imageTemplate *v1alpha1.ImageTemplate, opts v1.UpdateOptions) (*v1alpha1.ImageTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(imagetemplatesResource, "status", c.ns, imageTemplate), &v1alpha1.ImageTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageTemplate), err
}

// Delete takes name of the imageTemplate and deletes it. Returns an error if one occurs.
func (c *FakeImageTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(imagetemplatesResource, c.ns, name), &v1alpha1.ImageTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImageTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(imagetemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ImageTemplateList{})
	return err
}

// Patch applies the patch and returns the patched imageTemplate.
func (c *FakeImageTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImageTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(imagetemplatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ImageTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageTemplate), err
}
