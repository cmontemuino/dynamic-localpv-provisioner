/*
Copyright 2019 The OpenEBS Authors

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

	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	scheme "github.com/openebs/maya/pkg/client/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CASTemplatesGetter has a method to return a CASTemplateInterface.
// A group's client should implement this interface.
type CASTemplatesGetter interface {
	CASTemplates() CASTemplateInterface
}

// CASTemplateInterface has methods to work with CASTemplate resources.
type CASTemplateInterface interface {
	Create(ctx context.Context, cASTemplate *v1alpha1.CASTemplate, opts v1.CreateOptions) (*v1alpha1.CASTemplate, error)
	Update(ctx context.Context, cASTemplate *v1alpha1.CASTemplate, opts v1.UpdateOptions) (*v1alpha1.CASTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CASTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CASTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CASTemplate, err error)
	CASTemplateExpansion
}

// cASTemplates implements CASTemplateInterface
type cASTemplates struct {
	client rest.Interface
}

// newCASTemplates returns a CASTemplates
func newCASTemplates(c *OpenebsV1alpha1Client) *cASTemplates {
	return &cASTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the cASTemplate, and returns the corresponding cASTemplate object, and an error if there is any.
func (c *cASTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CASTemplate, err error) {
	result = &v1alpha1.CASTemplate{}
	err = c.client.Get().
		Resource("castemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CASTemplates that match those selectors.
func (c *cASTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CASTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CASTemplateList{}
	err = c.client.Get().
		Resource("castemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cASTemplates.
func (c *cASTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("castemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cASTemplate and creates it.  Returns the server's representation of the cASTemplate, and an error, if there is any.
func (c *cASTemplates) Create(ctx context.Context, cASTemplate *v1alpha1.CASTemplate, opts v1.CreateOptions) (result *v1alpha1.CASTemplate, err error) {
	result = &v1alpha1.CASTemplate{}
	err = c.client.Post().
		Resource("castemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cASTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cASTemplate and updates it. Returns the server's representation of the cASTemplate, and an error, if there is any.
func (c *cASTemplates) Update(ctx context.Context, cASTemplate *v1alpha1.CASTemplate, opts v1.UpdateOptions) (result *v1alpha1.CASTemplate, err error) {
	result = &v1alpha1.CASTemplate{}
	err = c.client.Put().
		Resource("castemplates").
		Name(cASTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cASTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cASTemplate and deletes it. Returns an error if one occurs.
func (c *cASTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("castemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cASTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("castemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cASTemplate.
func (c *cASTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CASTemplate, err error) {
	result = &v1alpha1.CASTemplate{}
	err = c.client.Patch(pt).
		Resource("castemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}