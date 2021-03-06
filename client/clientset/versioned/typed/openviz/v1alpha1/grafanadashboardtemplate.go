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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "go.openviz.dev/apimachinery/apis/openviz/v1alpha1"
	scheme "go.openviz.dev/apimachinery/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GrafanaDashboardTemplatesGetter has a method to return a GrafanaDashboardTemplateInterface.
// A group's client should implement this interface.
type GrafanaDashboardTemplatesGetter interface {
	GrafanaDashboardTemplates(namespace string) GrafanaDashboardTemplateInterface
}

// GrafanaDashboardTemplateInterface has methods to work with GrafanaDashboardTemplate resources.
type GrafanaDashboardTemplateInterface interface {
	Create(ctx context.Context, grafanaDashboardTemplate *v1alpha1.GrafanaDashboardTemplate, opts v1.CreateOptions) (*v1alpha1.GrafanaDashboardTemplate, error)
	Update(ctx context.Context, grafanaDashboardTemplate *v1alpha1.GrafanaDashboardTemplate, opts v1.UpdateOptions) (*v1alpha1.GrafanaDashboardTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.GrafanaDashboardTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.GrafanaDashboardTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GrafanaDashboardTemplate, err error)
	GrafanaDashboardTemplateExpansion
}

// grafanaDashboardTemplates implements GrafanaDashboardTemplateInterface
type grafanaDashboardTemplates struct {
	client rest.Interface
	ns     string
}

// newGrafanaDashboardTemplates returns a GrafanaDashboardTemplates
func newGrafanaDashboardTemplates(c *OpenvizV1alpha1Client, namespace string) *grafanaDashboardTemplates {
	return &grafanaDashboardTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the grafanaDashboardTemplate, and returns the corresponding grafanaDashboardTemplate object, and an error if there is any.
func (c *grafanaDashboardTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GrafanaDashboardTemplate, err error) {
	result = &v1alpha1.GrafanaDashboardTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("grafanadashboardtemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GrafanaDashboardTemplates that match those selectors.
func (c *grafanaDashboardTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GrafanaDashboardTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GrafanaDashboardTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("grafanadashboardtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested grafanaDashboardTemplates.
func (c *grafanaDashboardTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("grafanadashboardtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a grafanaDashboardTemplate and creates it.  Returns the server's representation of the grafanaDashboardTemplate, and an error, if there is any.
func (c *grafanaDashboardTemplates) Create(ctx context.Context, grafanaDashboardTemplate *v1alpha1.GrafanaDashboardTemplate, opts v1.CreateOptions) (result *v1alpha1.GrafanaDashboardTemplate, err error) {
	result = &v1alpha1.GrafanaDashboardTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("grafanadashboardtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(grafanaDashboardTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a grafanaDashboardTemplate and updates it. Returns the server's representation of the grafanaDashboardTemplate, and an error, if there is any.
func (c *grafanaDashboardTemplates) Update(ctx context.Context, grafanaDashboardTemplate *v1alpha1.GrafanaDashboardTemplate, opts v1.UpdateOptions) (result *v1alpha1.GrafanaDashboardTemplate, err error) {
	result = &v1alpha1.GrafanaDashboardTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("grafanadashboardtemplates").
		Name(grafanaDashboardTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(grafanaDashboardTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the grafanaDashboardTemplate and deletes it. Returns an error if one occurs.
func (c *grafanaDashboardTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("grafanadashboardtemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *grafanaDashboardTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("grafanadashboardtemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched grafanaDashboardTemplate.
func (c *grafanaDashboardTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GrafanaDashboardTemplate, err error) {
	result = &v1alpha1.GrafanaDashboardTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("grafanadashboardtemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
