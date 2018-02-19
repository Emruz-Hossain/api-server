/*
  Copyright 2017 The Kubernetes Authors.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/package fake

import (
	customrc "api-server/pkg/apis/customrc"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCustomReplicationControllers implements CustomReplicationControllerInterface
type FakeCustomReplicationControllers struct {
	Fake *FakeCustomrc
	ns   string
}

var customreplicationcontrollersResource = schema.GroupVersionResource{Group: "customrc.emruz.com", Version: "", Resource: "customreplicationcontrollers"}

var customreplicationcontrollersKind = schema.GroupVersionKind{Group: "customrc.emruz.com", Version: "", Kind: "CustomReplicationController"}

// Get takes name of the customReplicationController, and returns the corresponding customReplicationController object, and an error if there is any.
func (c *FakeCustomReplicationControllers) Get(name string, options v1.GetOptions) (result *customrc.CustomReplicationController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(customreplicationcontrollersResource, c.ns, name), &customrc.CustomReplicationController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*customrc.CustomReplicationController), err
}

// List takes label and field selectors, and returns the list of CustomReplicationControllers that match those selectors.
func (c *FakeCustomReplicationControllers) List(opts v1.ListOptions) (result *customrc.CustomReplicationControllerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(customreplicationcontrollersResource, customreplicationcontrollersKind, c.ns, opts), &customrc.CustomReplicationControllerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &customrc.CustomReplicationControllerList{}
	for _, item := range obj.(*customrc.CustomReplicationControllerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested customReplicationControllers.
func (c *FakeCustomReplicationControllers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(customreplicationcontrollersResource, c.ns, opts))

}

// Create takes the representation of a customReplicationController and creates it.  Returns the server's representation of the customReplicationController, and an error, if there is any.
func (c *FakeCustomReplicationControllers) Create(customReplicationController *customrc.CustomReplicationController) (result *customrc.CustomReplicationController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(customreplicationcontrollersResource, c.ns, customReplicationController), &customrc.CustomReplicationController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*customrc.CustomReplicationController), err
}

// Update takes the representation of a customReplicationController and updates it. Returns the server's representation of the customReplicationController, and an error, if there is any.
func (c *FakeCustomReplicationControllers) Update(customReplicationController *customrc.CustomReplicationController) (result *customrc.CustomReplicationController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(customreplicationcontrollersResource, c.ns, customReplicationController), &customrc.CustomReplicationController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*customrc.CustomReplicationController), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCustomReplicationControllers) UpdateStatus(customReplicationController *customrc.CustomReplicationController) (*customrc.CustomReplicationController, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(customreplicationcontrollersResource, "status", c.ns, customReplicationController), &customrc.CustomReplicationController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*customrc.CustomReplicationController), err
}

// Delete takes name of the customReplicationController and deletes it. Returns an error if one occurs.
func (c *FakeCustomReplicationControllers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(customreplicationcontrollersResource, c.ns, name), &customrc.CustomReplicationController{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCustomReplicationControllers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(customreplicationcontrollersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &customrc.CustomReplicationControllerList{})
	return err
}

// Patch applies the patch and returns the patched customReplicationController.
func (c *FakeCustomReplicationControllers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *customrc.CustomReplicationController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(customreplicationcontrollersResource, c.ns, name, data, subresources...), &customrc.CustomReplicationController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*customrc.CustomReplicationController), err
}
