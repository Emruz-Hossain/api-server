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
	clientset "api-server/pkg/client/clientset_generated/clientset"
	customdpv1alpha1 "api-server/pkg/client/clientset_generated/clientset/typed/customdp/v1alpha1"
	fakecustomdpv1alpha1 "api-server/pkg/client/clientset_generated/clientset/typed/customdp/v1alpha1/fake"
	customrcv1alpha1 "api-server/pkg/client/clientset_generated/clientset/typed/customrc/v1alpha1"
	fakecustomrcv1alpha1 "api-server/pkg/client/clientset_generated/clientset/typed/customrc/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o))
	fakePtr.AddWatchReactor("*", testing.DefaultWatchReactor(watch.NewFake(), nil))

	return &Clientset{fakePtr, &fakediscovery.FakeDiscovery{Fake: &fakePtr}}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// CustomdpV1alpha1 retrieves the CustomdpV1alpha1Client
func (c *Clientset) CustomdpV1alpha1() customdpv1alpha1.CustomdpV1alpha1Interface {
	return &fakecustomdpv1alpha1.FakeCustomdpV1alpha1{Fake: &c.Fake}
}

// Customdp retrieves the CustomdpV1alpha1Client
func (c *Clientset) Customdp() customdpv1alpha1.CustomdpV1alpha1Interface {
	return &fakecustomdpv1alpha1.FakeCustomdpV1alpha1{Fake: &c.Fake}
}

// CustomrcV1alpha1 retrieves the CustomrcV1alpha1Client
func (c *Clientset) CustomrcV1alpha1() customrcv1alpha1.CustomrcV1alpha1Interface {
	return &fakecustomrcv1alpha1.FakeCustomrcV1alpha1{Fake: &c.Fake}
}

// Customrc retrieves the CustomrcV1alpha1Client
func (c *Clientset) Customrc() customrcv1alpha1.CustomrcV1alpha1Interface {
	return &fakecustomrcv1alpha1.FakeCustomrcV1alpha1{Fake: &c.Fake}
}
