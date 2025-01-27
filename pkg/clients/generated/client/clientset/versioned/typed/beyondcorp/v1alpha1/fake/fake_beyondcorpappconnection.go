// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/beyondcorp/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBeyondCorpAppConnections implements BeyondCorpAppConnectionInterface
type FakeBeyondCorpAppConnections struct {
	Fake *FakeBeyondcorpV1alpha1
	ns   string
}

var beyondcorpappconnectionsResource = schema.GroupVersionResource{Group: "beyondcorp.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "beyondcorpappconnections"}

var beyondcorpappconnectionsKind = schema.GroupVersionKind{Group: "beyondcorp.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BeyondCorpAppConnection"}

// Get takes name of the beyondCorpAppConnection, and returns the corresponding beyondCorpAppConnection object, and an error if there is any.
func (c *FakeBeyondCorpAppConnections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BeyondCorpAppConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(beyondcorpappconnectionsResource, c.ns, name), &v1alpha1.BeyondCorpAppConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BeyondCorpAppConnection), err
}

// List takes label and field selectors, and returns the list of BeyondCorpAppConnections that match those selectors.
func (c *FakeBeyondCorpAppConnections) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BeyondCorpAppConnectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(beyondcorpappconnectionsResource, beyondcorpappconnectionsKind, c.ns, opts), &v1alpha1.BeyondCorpAppConnectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BeyondCorpAppConnectionList{ListMeta: obj.(*v1alpha1.BeyondCorpAppConnectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.BeyondCorpAppConnectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested beyondCorpAppConnections.
func (c *FakeBeyondCorpAppConnections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(beyondcorpappconnectionsResource, c.ns, opts))

}

// Create takes the representation of a beyondCorpAppConnection and creates it.  Returns the server's representation of the beyondCorpAppConnection, and an error, if there is any.
func (c *FakeBeyondCorpAppConnections) Create(ctx context.Context, beyondCorpAppConnection *v1alpha1.BeyondCorpAppConnection, opts v1.CreateOptions) (result *v1alpha1.BeyondCorpAppConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(beyondcorpappconnectionsResource, c.ns, beyondCorpAppConnection), &v1alpha1.BeyondCorpAppConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BeyondCorpAppConnection), err
}

// Update takes the representation of a beyondCorpAppConnection and updates it. Returns the server's representation of the beyondCorpAppConnection, and an error, if there is any.
func (c *FakeBeyondCorpAppConnections) Update(ctx context.Context, beyondCorpAppConnection *v1alpha1.BeyondCorpAppConnection, opts v1.UpdateOptions) (result *v1alpha1.BeyondCorpAppConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(beyondcorpappconnectionsResource, c.ns, beyondCorpAppConnection), &v1alpha1.BeyondCorpAppConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BeyondCorpAppConnection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBeyondCorpAppConnections) UpdateStatus(ctx context.Context, beyondCorpAppConnection *v1alpha1.BeyondCorpAppConnection, opts v1.UpdateOptions) (*v1alpha1.BeyondCorpAppConnection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(beyondcorpappconnectionsResource, "status", c.ns, beyondCorpAppConnection), &v1alpha1.BeyondCorpAppConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BeyondCorpAppConnection), err
}

// Delete takes name of the beyondCorpAppConnection and deletes it. Returns an error if one occurs.
func (c *FakeBeyondCorpAppConnections) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(beyondcorpappconnectionsResource, c.ns, name, opts), &v1alpha1.BeyondCorpAppConnection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBeyondCorpAppConnections) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(beyondcorpappconnectionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BeyondCorpAppConnectionList{})
	return err
}

// Patch applies the patch and returns the patched beyondCorpAppConnection.
func (c *FakeBeyondCorpAppConnections) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BeyondCorpAppConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(beyondcorpappconnectionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BeyondCorpAppConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BeyondCorpAppConnection), err
}
