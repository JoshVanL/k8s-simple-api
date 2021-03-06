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

package fake

import (
	v1alpha1 "github.com/joshvanl/k8s-simple-api/pkg/apis/simple/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMessages implements MessageInterface
type FakeMessages struct {
	Fake *FakeSimpleV1alpha1
	ns   string
}

var messagesResource = schema.GroupVersionResource{Group: "simple", Version: "v1alpha1", Resource: "messages"}

var messagesKind = schema.GroupVersionKind{Group: "simple", Version: "v1alpha1", Kind: "Message"}

func (c *FakeMessages) Create(message *v1alpha1.Message) (result *v1alpha1.Message, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(messagesResource, c.ns, message), &v1alpha1.Message{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Message), err
}

func (c *FakeMessages) Update(message *v1alpha1.Message) (result *v1alpha1.Message, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(messagesResource, c.ns, message), &v1alpha1.Message{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Message), err
}

func (c *FakeMessages) UpdateStatus(message *v1alpha1.Message) (*v1alpha1.Message, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(messagesResource, "status", c.ns, message), &v1alpha1.Message{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Message), err
}

func (c *FakeMessages) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(messagesResource, c.ns, name), &v1alpha1.Message{})

	return err
}

func (c *FakeMessages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(messagesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MessageList{})
	return err
}

func (c *FakeMessages) Get(name string, options v1.GetOptions) (result *v1alpha1.Message, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(messagesResource, c.ns, name), &v1alpha1.Message{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Message), err
}

func (c *FakeMessages) List(opts v1.ListOptions) (result *v1alpha1.MessageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(messagesResource, messagesKind, c.ns, opts), &v1alpha1.MessageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MessageList{}
	for _, item := range obj.(*v1alpha1.MessageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested messages.
func (c *FakeMessages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(messagesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched message.
func (c *FakeMessages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Message, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(messagesResource, c.ns, name, data, subresources...), &v1alpha1.Message{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Message), err
}
