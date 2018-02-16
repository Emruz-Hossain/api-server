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
*/
// This file was automatically generated by informer-gen

package v1alpha1

import (
	car_v1alpha1 "api-server/pkg/apis/car/v1alpha1"
	clientset "api-server/pkg/client/clientset_generated/clientset"
	internalinterfaces "api-server/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "api-server/pkg/client/listers_generated/car/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// CarControllerInformer provides access to a shared informer and lister for
// CarControllers.
type CarControllerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CarControllerLister
}

type carControllerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCarControllerInformer constructs a new informer for CarController type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCarControllerInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCarControllerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCarControllerInformer constructs a new informer for CarController type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCarControllerInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CarV1alpha1().CarControllers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CarV1alpha1().CarControllers(namespace).Watch(options)
			},
		},
		&car_v1alpha1.CarController{},
		resyncPeriod,
		indexers,
	)
}

func (f *carControllerInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCarControllerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *carControllerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&car_v1alpha1.CarController{}, f.defaultInformer)
}

func (f *carControllerInformer) Lister() v1alpha1.CarControllerLister {
	return v1alpha1.NewCarControllerLister(f.Informer().GetIndexer())
}