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
	customdp "api-server/pkg/apis/customdp"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCustomDeployments implements CustomDeploymentInterface
type FakeCustomDeployments struct {
	Fake *FakeCustomdp
	ns   string
}

var customdeploymentsResource = schema.GroupVersionResource{Group: "customdp.emruz.com", Version: "", Resource: "customdeployments"}

var customdeploymentsKind = schema.GroupVersionKind{Group: "customdp.emruz.com", Version: "", Kind: "CustomDeployment"}

// Get takes name of the customDeployment, and returns the corresponding customDeployment object, and an error if there is any.
func (c *FakeCustomDeployments) Get(name string, options v1.GetOptions) (result *customdp.CustomDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(customdeploymentsResource, c.ns, name), &customdp.CustomDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*customdp.CustomDeployment), err
}

// List takes label and field selectors, and returns the list of CustomDeployments that match those selectors.
func (c *FakeCustomDeployments) List(opts v1.ListOptions) (result *customdp.CustomDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(customdeploymentsResource, customdeploymentsKind, c.ns, opts), &customdp.CustomDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &customdp.CustomDeploymentList{}
	for _, item := range obj.(*customdp.CustomDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested customDeployments.
func (c *FakeCustomDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(customdeploymentsResource, c.ns, opts))

}

// Create takes the representation of a customDeployment and creates it.  Returns the server's representation of the customDeployment, and an error, if there is any.
func (c *FakeCustomDeployments) Create(customDeployment *customdp.CustomDeployment) (result *customdp.CustomDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(customdeploymentsResource, c.ns, customDeployment), &customdp.CustomDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*customdp.CustomDeployment), err
}

// Update takes the representation of a customDeployment and updates it. Returns the server's representation of the customDeployment, and an error, if there is any.
func (c *FakeCustomDeployments) Update(customDeployment *customdp.CustomDeployment) (result *customdp.CustomDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(customdeploymentsResource, c.ns, customDeployment), &customdp.CustomDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*customdp.CustomDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCustomDeployments) UpdateStatus(customDeployment *customdp.CustomDeployment) (*customdp.CustomDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(customdeploymentsResource, "status", c.ns, customDeployment), &customdp.CustomDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*customdp.CustomDeployment), err
}

// Delete takes name of the customDeployment and deletes it. Returns an error if one occurs.
func (c *FakeCustomDeployments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(customdeploymentsResource, c.ns, name), &customdp.CustomDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCustomDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(customdeploymentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &customdp.CustomDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched customDeployment.
func (c *FakeCustomDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *customdp.CustomDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(customdeploymentsResource, c.ns, name, data, subresources...), &customdp.CustomDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*customdp.CustomDeployment), err
}
