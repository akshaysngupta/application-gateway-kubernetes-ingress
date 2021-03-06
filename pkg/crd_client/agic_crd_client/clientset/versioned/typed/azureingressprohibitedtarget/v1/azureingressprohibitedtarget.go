/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"time"

	v1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/apis/azureingressprohibitedtarget/v1"
	scheme "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/agic_crd_client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AzureIngressProhibitedTargetsGetter has a method to return a AzureIngressProhibitedTargetInterface.
// A group's client should implement this interface.
type AzureIngressProhibitedTargetsGetter interface {
	AzureIngressProhibitedTargets(namespace string) AzureIngressProhibitedTargetInterface
}

// AzureIngressProhibitedTargetInterface has methods to work with AzureIngressProhibitedTarget resources.
type AzureIngressProhibitedTargetInterface interface {
	Create(*v1.AzureIngressProhibitedTarget) (*v1.AzureIngressProhibitedTarget, error)
	Update(*v1.AzureIngressProhibitedTarget) (*v1.AzureIngressProhibitedTarget, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.AzureIngressProhibitedTarget, error)
	List(opts metav1.ListOptions) (*v1.AzureIngressProhibitedTargetList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AzureIngressProhibitedTarget, err error)
	AzureIngressProhibitedTargetExpansion
}

// azureIngressProhibitedTargets implements AzureIngressProhibitedTargetInterface
type azureIngressProhibitedTargets struct {
	client rest.Interface
	ns     string
}

// newAzureIngressProhibitedTargets returns a AzureIngressProhibitedTargets
func newAzureIngressProhibitedTargets(c *AzureingressprohibitedtargetsV1Client, namespace string) *azureIngressProhibitedTargets {
	return &azureIngressProhibitedTargets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the azureIngressProhibitedTarget, and returns the corresponding azureIngressProhibitedTarget object, and an error if there is any.
func (c *azureIngressProhibitedTargets) Get(name string, options metav1.GetOptions) (result *v1.AzureIngressProhibitedTarget, err error) {
	result = &v1.AzureIngressProhibitedTarget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureingressprohibitedtargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzureIngressProhibitedTargets that match those selectors.
func (c *azureIngressProhibitedTargets) List(opts metav1.ListOptions) (result *v1.AzureIngressProhibitedTargetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AzureIngressProhibitedTargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureingressprohibitedtargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azureIngressProhibitedTargets.
func (c *azureIngressProhibitedTargets) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("azureingressprohibitedtargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azureIngressProhibitedTarget and creates it.  Returns the server's representation of the azureIngressProhibitedTarget, and an error, if there is any.
func (c *azureIngressProhibitedTargets) Create(azureIngressProhibitedTarget *v1.AzureIngressProhibitedTarget) (result *v1.AzureIngressProhibitedTarget, err error) {
	result = &v1.AzureIngressProhibitedTarget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("azureingressprohibitedtargets").
		Body(azureIngressProhibitedTarget).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azureIngressProhibitedTarget and updates it. Returns the server's representation of the azureIngressProhibitedTarget, and an error, if there is any.
func (c *azureIngressProhibitedTargets) Update(azureIngressProhibitedTarget *v1.AzureIngressProhibitedTarget) (result *v1.AzureIngressProhibitedTarget, err error) {
	result = &v1.AzureIngressProhibitedTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azureingressprohibitedtargets").
		Name(azureIngressProhibitedTarget.Name).
		Body(azureIngressProhibitedTarget).
		Do().
		Into(result)
	return
}

// Delete takes name of the azureIngressProhibitedTarget and deletes it. Returns an error if one occurs.
func (c *azureIngressProhibitedTargets) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureingressprohibitedtargets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azureIngressProhibitedTargets) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureingressprohibitedtargets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azureIngressProhibitedTarget.
func (c *azureIngressProhibitedTargets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AzureIngressProhibitedTarget, err error) {
	result = &v1.AzureIngressProhibitedTarget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("azureingressprohibitedtargets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
