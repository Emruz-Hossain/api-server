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
*/package internalversion

import (
	customrc "api-server/pkg/apis/customrc"
	scheme "api-server/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CustomReplicationControllersGetter has a method to return a CustomReplicationControllerInterface.
// A group's client should implement this interface.
type CustomReplicationControllersGetter interface {
	CustomReplicationControllers(namespace string) CustomReplicationControllerInterface
}

// CustomReplicationControllerInterface has methods to work with CustomReplicationController resources.
type CustomReplicationControllerInterface interface {
	Create(*customrc.CustomReplicationController) (*customrc.CustomReplicationController, error)
	Update(*customrc.CustomReplicationController) (*customrc.CustomReplicationController, error)
	UpdateStatus(*customrc.CustomReplicationController) (*customrc.CustomReplicationController, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*customrc.CustomReplicationController, error)
	List(opts v1.ListOptions) (*customrc.CustomReplicationControllerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *customrc.CustomReplicationController, err error)
	CustomReplicationControllerExpansion
}

// customReplicationControllers implements CustomReplicationControllerInterface
type customReplicationControllers struct {
	client rest.Interface
	ns     string
}

// newCustomReplicationControllers returns a CustomReplicationControllers
func newCustomReplicationControllers(c *CustomrcClient, namespace string) *customReplicationControllers {
	return &customReplicationControllers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the customReplicationController, and returns the corresponding customReplicationController object, and an error if there is any.
func (c *customReplicationControllers) Get(name string, options v1.GetOptions) (result *customrc.CustomReplicationController, err error) {
	result = &customrc.CustomReplicationController{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("customreplicationcontrollers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CustomReplicationControllers that match those selectors.
func (c *customReplicationControllers) List(opts v1.ListOptions) (result *customrc.CustomReplicationControllerList, err error) {
	result = &customrc.CustomReplicationControllerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("customreplicationcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested customReplicationControllers.
func (c *customReplicationControllers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("customreplicationcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a customReplicationController and creates it.  Returns the server's representation of the customReplicationController, and an error, if there is any.
func (c *customReplicationControllers) Create(customReplicationController *customrc.CustomReplicationController) (result *customrc.CustomReplicationController, err error) {
	result = &customrc.CustomReplicationController{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("customreplicationcontrollers").
		Body(customReplicationController).
		Do().
		Into(result)
	return
}

// Update takes the representation of a customReplicationController and updates it. Returns the server's representation of the customReplicationController, and an error, if there is any.
func (c *customReplicationControllers) Update(customReplicationController *customrc.CustomReplicationController) (result *customrc.CustomReplicationController, err error) {
	result = &customrc.CustomReplicationController{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("customreplicationcontrollers").
		Name(customReplicationController.Name).
		Body(customReplicationController).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *customReplicationControllers) UpdateStatus(customReplicationController *customrc.CustomReplicationController) (result *customrc.CustomReplicationController, err error) {
	result = &customrc.CustomReplicationController{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("customreplicationcontrollers").
		Name(customReplicationController.Name).
		SubResource("status").
		Body(customReplicationController).
		Do().
		Into(result)
	return
}

// Delete takes name of the customReplicationController and deletes it. Returns an error if one occurs.
func (c *customReplicationControllers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("customreplicationcontrollers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *customReplicationControllers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("customreplicationcontrollers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched customReplicationController.
func (c *customReplicationControllers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *customrc.CustomReplicationController, err error) {
	result = &customrc.CustomReplicationController{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("customreplicationcontrollers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
