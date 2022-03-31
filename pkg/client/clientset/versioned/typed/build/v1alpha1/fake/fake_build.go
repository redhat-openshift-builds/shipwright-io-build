// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBuilds implements BuildInterface
type FakeBuilds struct {
	Fake *FakeShipwrightV1alpha1
	ns   string
}

var buildsResource = schema.GroupVersionResource{Group: "shipwright.io", Version: "v1alpha1", Resource: "builds"}

var buildsKind = schema.GroupVersionKind{Group: "shipwright.io", Version: "v1alpha1", Kind: "Build"}

// Get takes name of the build, and returns the corresponding build object, and an error if there is any.
func (c *FakeBuilds) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(buildsResource, c.ns, name), &v1alpha1.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Build), err
}

// List takes label and field selectors, and returns the list of Builds that match those selectors.
func (c *FakeBuilds) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BuildList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(buildsResource, buildsKind, c.ns, opts), &v1alpha1.BuildList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BuildList{ListMeta: obj.(*v1alpha1.BuildList).ListMeta}
	for _, item := range obj.(*v1alpha1.BuildList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested builds.
func (c *FakeBuilds) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(buildsResource, c.ns, opts))

}

// Create takes the representation of a build and creates it.  Returns the server's representation of the build, and an error, if there is any.
func (c *FakeBuilds) Create(ctx context.Context, build *v1alpha1.Build, opts v1.CreateOptions) (result *v1alpha1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(buildsResource, c.ns, build), &v1alpha1.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Build), err
}

// Update takes the representation of a build and updates it. Returns the server's representation of the build, and an error, if there is any.
func (c *FakeBuilds) Update(ctx context.Context, build *v1alpha1.Build, opts v1.UpdateOptions) (result *v1alpha1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(buildsResource, c.ns, build), &v1alpha1.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Build), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBuilds) UpdateStatus(ctx context.Context, build *v1alpha1.Build, opts v1.UpdateOptions) (*v1alpha1.Build, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(buildsResource, "status", c.ns, build), &v1alpha1.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Build), err
}

// Delete takes name of the build and deletes it. Returns an error if one occurs.
func (c *FakeBuilds) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(buildsResource, c.ns, name, opts), &v1alpha1.Build{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuilds) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(buildsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BuildList{})
	return err
}

// Patch applies the patch and returns the patched build.
func (c *FakeBuilds) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Build), err
}
