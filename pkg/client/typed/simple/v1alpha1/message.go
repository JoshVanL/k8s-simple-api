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

package v1alpha1

import (
	v1alpha1 "github.com/joshvanl/k8s-simple-api/pkg/apis/simple/v1alpha1"
	scheme "github.com/joshvanl/k8s-simple-api/pkg/client/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MessagesGetter has a method to return a MessageInterface.
// A group's client should implement this interface.
type MessagesGetter interface {
	Messages(namespace string) MessageInterface
}

// MessageInterface has methods to work with Message resources.
type MessageInterface interface {
	Create(*v1alpha1.Message) (*v1alpha1.Message, error)
	Update(*v1alpha1.Message) (*v1alpha1.Message, error)
	UpdateStatus(*v1alpha1.Message) (*v1alpha1.Message, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Message, error)
	List(opts v1.ListOptions) (*v1alpha1.MessageList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Message, err error)
	MessageExpansion
}

// messages implements MessageInterface
type messages struct {
	client rest.Interface
	ns     string
}

// newMessages returns a Messages
func newMessages(c *SimpleV1alpha1Client, namespace string) *messages {
	return &messages{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a message and creates it.  Returns the server's representation of the message, and an error, if there is any.
func (c *messages) Create(message *v1alpha1.Message) (result *v1alpha1.Message, err error) {
	result = &v1alpha1.Message{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("messages").
		Body(message).
		Do().
		Into(result)
	return
}

// Update takes the representation of a message and updates it. Returns the server's representation of the message, and an error, if there is any.
func (c *messages) Update(message *v1alpha1.Message) (result *v1alpha1.Message, err error) {
	result = &v1alpha1.Message{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("messages").
		Name(message.Name).
		Body(message).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclientstatus=false comment above the type to avoid generating UpdateStatus().

func (c *messages) UpdateStatus(message *v1alpha1.Message) (result *v1alpha1.Message, err error) {
	result = &v1alpha1.Message{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("messages").
		Name(message.Name).
		SubResource("status").
		Body(message).
		Do().
		Into(result)
	return
}

// Delete takes name of the message and deletes it. Returns an error if one occurs.
func (c *messages) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("messages").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *messages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("messages").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the message, and returns the corresponding message object, and an error if there is any.
func (c *messages) Get(name string, options v1.GetOptions) (result *v1alpha1.Message, err error) {
	result = &v1alpha1.Message{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("messages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Messages that match those selectors.
func (c *messages) List(opts v1.ListOptions) (result *v1alpha1.MessageList, err error) {
	result = &v1alpha1.MessageList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("messages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested messages.
func (c *messages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("messages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched message.
func (c *messages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Message, err error) {
	result = &v1alpha1.Message{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("messages").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
