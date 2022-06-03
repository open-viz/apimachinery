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

package fake

import (
	"context"

	v1alpha1 "go.openviz.dev/apimachinery/apis/openviz/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGrafanaDatasources implements GrafanaDatasourceInterface
type FakeGrafanaDatasources struct {
	Fake *FakeOpenvizV1alpha1
	ns   string
}

var grafanadatasourcesResource = schema.GroupVersionResource{Group: "openviz.dev", Version: "v1alpha1", Resource: "grafanadatasources"}

var grafanadatasourcesKind = schema.GroupVersionKind{Group: "openviz.dev", Version: "v1alpha1", Kind: "GrafanaDatasource"}

// Get takes name of the grafanaDatasource, and returns the corresponding grafanaDatasource object, and an error if there is any.
func (c *FakeGrafanaDatasources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GrafanaDatasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(grafanadatasourcesResource, c.ns, name), &v1alpha1.GrafanaDatasource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaDatasource), err
}

// List takes label and field selectors, and returns the list of GrafanaDatasources that match those selectors.
func (c *FakeGrafanaDatasources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GrafanaDatasourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(grafanadatasourcesResource, grafanadatasourcesKind, c.ns, opts), &v1alpha1.GrafanaDatasourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GrafanaDatasourceList{ListMeta: obj.(*v1alpha1.GrafanaDatasourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.GrafanaDatasourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested grafanaDatasources.
func (c *FakeGrafanaDatasources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(grafanadatasourcesResource, c.ns, opts))

}

// Create takes the representation of a grafanaDatasource and creates it.  Returns the server's representation of the grafanaDatasource, and an error, if there is any.
func (c *FakeGrafanaDatasources) Create(ctx context.Context, grafanaDatasource *v1alpha1.GrafanaDatasource, opts v1.CreateOptions) (result *v1alpha1.GrafanaDatasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(grafanadatasourcesResource, c.ns, grafanaDatasource), &v1alpha1.GrafanaDatasource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaDatasource), err
}

// Update takes the representation of a grafanaDatasource and updates it. Returns the server's representation of the grafanaDatasource, and an error, if there is any.
func (c *FakeGrafanaDatasources) Update(ctx context.Context, grafanaDatasource *v1alpha1.GrafanaDatasource, opts v1.UpdateOptions) (result *v1alpha1.GrafanaDatasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(grafanadatasourcesResource, c.ns, grafanaDatasource), &v1alpha1.GrafanaDatasource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaDatasource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGrafanaDatasources) UpdateStatus(ctx context.Context, grafanaDatasource *v1alpha1.GrafanaDatasource, opts v1.UpdateOptions) (*v1alpha1.GrafanaDatasource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(grafanadatasourcesResource, "status", c.ns, grafanaDatasource), &v1alpha1.GrafanaDatasource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaDatasource), err
}

// Delete takes name of the grafanaDatasource and deletes it. Returns an error if one occurs.
func (c *FakeGrafanaDatasources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(grafanadatasourcesResource, c.ns, name, opts), &v1alpha1.GrafanaDatasource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGrafanaDatasources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(grafanadatasourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GrafanaDatasourceList{})
	return err
}

// Patch applies the patch and returns the patched grafanaDatasource.
func (c *FakeGrafanaDatasources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GrafanaDatasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(grafanadatasourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.GrafanaDatasource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaDatasource), err
}
