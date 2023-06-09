/*
Copyright 2023 The KServe Authors.

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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/kserve/kserve/pkg/apis/serving/v1beta1"
	scheme "github.com/kserve/kserve/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InferenceServicesGetter has a method to return a InferenceServiceInterface.
// A group's client should implement this interface.
type InferenceServicesGetter interface {
	InferenceServices(namespace string) InferenceServiceInterface
}

// InferenceServiceInterface has methods to work with InferenceService resources.
type InferenceServiceInterface interface {
	Create(ctx context.Context, inferenceService *v1beta1.InferenceService, opts v1.CreateOptions) (*v1beta1.InferenceService, error)
	Update(ctx context.Context, inferenceService *v1beta1.InferenceService, opts v1.UpdateOptions) (*v1beta1.InferenceService, error)
	UpdateStatus(ctx context.Context, inferenceService *v1beta1.InferenceService, opts v1.UpdateOptions) (*v1beta1.InferenceService, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.InferenceService, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.InferenceServiceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.InferenceService, err error)
	InferenceServiceExpansion
}

// inferenceServices implements InferenceServiceInterface
type inferenceServices struct {
	client rest.Interface
	ns     string
}

// newInferenceServices returns a InferenceServices
func newInferenceServices(c *ServingV1beta1Client, namespace string) *inferenceServices {
	return &inferenceServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the inferenceService, and returns the corresponding inferenceService object, and an error if there is any.
func (c *inferenceServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.InferenceService, err error) {
	result = &v1beta1.InferenceService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("inferenceservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of InferenceServices that match those selectors.
func (c *inferenceServices) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.InferenceServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.InferenceServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("inferenceservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested inferenceServices.
func (c *inferenceServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("inferenceservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a inferenceService and creates it.  Returns the server's representation of the inferenceService, and an error, if there is any.
func (c *inferenceServices) Create(ctx context.Context, inferenceService *v1beta1.InferenceService, opts v1.CreateOptions) (result *v1beta1.InferenceService, err error) {
	result = &v1beta1.InferenceService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("inferenceservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(inferenceService).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a inferenceService and updates it. Returns the server's representation of the inferenceService, and an error, if there is any.
func (c *inferenceServices) Update(ctx context.Context, inferenceService *v1beta1.InferenceService, opts v1.UpdateOptions) (result *v1beta1.InferenceService, err error) {
	result = &v1beta1.InferenceService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("inferenceservices").
		Name(inferenceService.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(inferenceService).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *inferenceServices) UpdateStatus(ctx context.Context, inferenceService *v1beta1.InferenceService, opts v1.UpdateOptions) (result *v1beta1.InferenceService, err error) {
	result = &v1beta1.InferenceService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("inferenceservices").
		Name(inferenceService.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(inferenceService).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the inferenceService and deletes it. Returns an error if one occurs.
func (c *inferenceServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("inferenceservices").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *inferenceServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("inferenceservices").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched inferenceService.
func (c *inferenceServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.InferenceService, err error) {
	result = &v1beta1.InferenceService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("inferenceservices").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
