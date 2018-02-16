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
*/package v1alpha1

import (
	v1alpha1 "api-server/pkg/apis/car/v1alpha1"
	scheme "api-server/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CarControllersGetter has a method to return a CarControllerInterface.
// A group's client should implement this interface.
type CarControllersGetter interface {
	CarControllers(namespace string) CarControllerInterface
}

// CarControllerInterface has methods to work with CarController resources.
type CarControllerInterface interface {
	Create(*v1alpha1.CarController) (*v1alpha1.CarController, error)
	Update(*v1alpha1.CarController) (*v1alpha1.CarController, error)
	UpdateStatus(*v1alpha1.CarController) (*v1alpha1.CarController, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CarController, error)
	List(opts v1.ListOptions) (*v1alpha1.CarControllerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CarController, err error)
	CarControllerExpansion
}

// carControllers implements CarControllerInterface
type carControllers struct {
	client rest.Interface
	ns     string
}

// newCarControllers returns a CarControllers
func newCarControllers(c *CarV1alpha1Client, namespace string) *carControllers {
	return &carControllers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the carController, and returns the corresponding carController object, and an error if there is any.
func (c *carControllers) Get(name string, options v1.GetOptions) (result *v1alpha1.CarController, err error) {
	result = &v1alpha1.CarController{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("carcontrollers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CarControllers that match those selectors.
func (c *carControllers) List(opts v1.ListOptions) (result *v1alpha1.CarControllerList, err error) {
	result = &v1alpha1.CarControllerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("carcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested carControllers.
func (c *carControllers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("carcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a carController and creates it.  Returns the server's representation of the carController, and an error, if there is any.
func (c *carControllers) Create(carController *v1alpha1.CarController) (result *v1alpha1.CarController, err error) {
	result = &v1alpha1.CarController{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("carcontrollers").
		Body(carController).
		Do().
		Into(result)
	return
}

// Update takes the representation of a carController and updates it. Returns the server's representation of the carController, and an error, if there is any.
func (c *carControllers) Update(carController *v1alpha1.CarController) (result *v1alpha1.CarController, err error) {
	result = &v1alpha1.CarController{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("carcontrollers").
		Name(carController.Name).
		Body(carController).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *carControllers) UpdateStatus(carController *v1alpha1.CarController) (result *v1alpha1.CarController, err error) {
	result = &v1alpha1.CarController{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("carcontrollers").
		Name(carController.Name).
		SubResource("status").
		Body(carController).
		Do().
		Into(result)
	return
}

// Delete takes name of the carController and deletes it. Returns an error if one occurs.
func (c *carControllers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("carcontrollers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *carControllers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("carcontrollers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched carController.
func (c *carControllers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CarController, err error) {
	result = &v1alpha1.CarController{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("carcontrollers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
