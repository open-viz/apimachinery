/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	openvizv1alpha1 "go.openviz.dev/grafana-tools/apis/openviz/v1alpha1"
	versioned "go.openviz.dev/grafana-tools/client/clientset/versioned"
	internalinterfaces "go.openviz.dev/grafana-tools/client/informers/externalversions/internalinterfaces"
	v1alpha1 "go.openviz.dev/grafana-tools/client/listers/openviz/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DashboardTemplateInformer provides access to a shared informer and lister for
// DashboardTemplates.
type DashboardTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DashboardTemplateLister
}

type dashboardTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDashboardTemplateInformer constructs a new informer for DashboardTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDashboardTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDashboardTemplateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDashboardTemplateInformer constructs a new informer for DashboardTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDashboardTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenvizV1alpha1().DashboardTemplates(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenvizV1alpha1().DashboardTemplates(namespace).Watch(context.TODO(), options)
			},
		},
		&openvizv1alpha1.DashboardTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *dashboardTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDashboardTemplateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dashboardTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&openvizv1alpha1.DashboardTemplate{}, f.defaultInformer)
}

func (f *dashboardTemplateInformer) Lister() v1alpha1.DashboardTemplateLister {
	return v1alpha1.NewDashboardTemplateLister(f.Informer().GetIndexer())
}
