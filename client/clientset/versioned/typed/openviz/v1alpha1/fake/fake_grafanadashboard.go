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

	v1alpha1 "go.openviz.dev/grafana-tools/apis/openviz/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGrafanaDashboards implements GrafanaDashboardInterface
type FakeGrafanaDashboards struct {
	Fake *FakeOpenvizV1alpha1
	ns   string
}

var grafanadashboardsResource = schema.GroupVersionResource{Group: "openviz.dev", Version: "v1alpha1", Resource: "grafanadashboards"}

var grafanadashboardsKind = schema.GroupVersionKind{Group: "openviz.dev", Version: "v1alpha1", Kind: "GrafanaDashboard"}

// Get takes name of the grafanaDashboard, and returns the corresponding grafanaDashboard object, and an error if there is any.
func (c *FakeGrafanaDashboards) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GrafanaDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(grafanadashboardsResource, c.ns, name), &v1alpha1.GrafanaDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaDashboard), err
}

// List takes label and field selectors, and returns the list of GrafanaDashboards that match those selectors.
func (c *FakeGrafanaDashboards) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GrafanaDashboardList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(grafanadashboardsResource, grafanadashboardsKind, c.ns, opts), &v1alpha1.GrafanaDashboardList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GrafanaDashboardList{ListMeta: obj.(*v1alpha1.GrafanaDashboardList).ListMeta}
	for _, item := range obj.(*v1alpha1.GrafanaDashboardList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested grafanaDashboards.
func (c *FakeGrafanaDashboards) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(grafanadashboardsResource, c.ns, opts))

}

// Create takes the representation of a grafanaDashboard and creates it.  Returns the server's representation of the grafanaDashboard, and an error, if there is any.
func (c *FakeGrafanaDashboards) Create(ctx context.Context, grafanaDashboard *v1alpha1.GrafanaDashboard, opts v1.CreateOptions) (result *v1alpha1.GrafanaDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(grafanadashboardsResource, c.ns, grafanaDashboard), &v1alpha1.GrafanaDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaDashboard), err
}

// Update takes the representation of a grafanaDashboard and updates it. Returns the server's representation of the grafanaDashboard, and an error, if there is any.
func (c *FakeGrafanaDashboards) Update(ctx context.Context, grafanaDashboard *v1alpha1.GrafanaDashboard, opts v1.UpdateOptions) (result *v1alpha1.GrafanaDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(grafanadashboardsResource, c.ns, grafanaDashboard), &v1alpha1.GrafanaDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaDashboard), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGrafanaDashboards) UpdateStatus(ctx context.Context, grafanaDashboard *v1alpha1.GrafanaDashboard, opts v1.UpdateOptions) (*v1alpha1.GrafanaDashboard, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(grafanadashboardsResource, "status", c.ns, grafanaDashboard), &v1alpha1.GrafanaDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaDashboard), err
}

// Delete takes name of the grafanaDashboard and deletes it. Returns an error if one occurs.
func (c *FakeGrafanaDashboards) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(grafanadashboardsResource, c.ns, name), &v1alpha1.GrafanaDashboard{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGrafanaDashboards) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(grafanadashboardsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GrafanaDashboardList{})
	return err
}

// Patch applies the patch and returns the patched grafanaDashboard.
func (c *FakeGrafanaDashboards) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GrafanaDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(grafanadashboardsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GrafanaDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaDashboard), err
}
