// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	privateregistryv1alpha1 "github.com/GLYASAI/rainbond-operator/pkg/apis/privateregistry/v1alpha1"
	versioned "github.com/GLYASAI/rainbond-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/GLYASAI/rainbond-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/GLYASAI/rainbond-operator/pkg/generated/listers/privateregistry/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PrivateRegistryInformer provides access to a shared informer and lister for
// PrivateRegistries.
type PrivateRegistryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PrivateRegistryLister
}

type privateRegistryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPrivateRegistryInformer constructs a new informer for PrivateRegistry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPrivateRegistryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPrivateRegistryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPrivateRegistryInformer constructs a new informer for PrivateRegistry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPrivateRegistryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PrivateregistryV1alpha1().PrivateRegistries(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PrivateregistryV1alpha1().PrivateRegistries(namespace).Watch(options)
			},
		},
		&privateregistryv1alpha1.PrivateRegistry{},
		resyncPeriod,
		indexers,
	)
}

func (f *privateRegistryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPrivateRegistryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *privateRegistryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&privateregistryv1alpha1.PrivateRegistry{}, f.defaultInformer)
}

func (f *privateRegistryInformer) Lister() v1alpha1.PrivateRegistryLister {
	return v1alpha1.NewPrivateRegistryLister(f.Informer().GetIndexer())
}
