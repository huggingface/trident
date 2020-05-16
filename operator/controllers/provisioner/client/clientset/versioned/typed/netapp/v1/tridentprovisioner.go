// Copyright 2020 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/netapp/trident/operator/controllers/provisioner/apis/netapp/v1"
	scheme "github.com/netapp/trident/operator/controllers/provisioner/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TridentProvisionersGetter has a method to return a TridentProvisionerInterface.
// A group's client should implement this interface.
type TridentProvisionersGetter interface {
	TridentProvisioners(namespace string) TridentProvisionerInterface
}

// TridentProvisionerInterface has methods to work with TridentProvisioner resources.
type TridentProvisionerInterface interface {
	Create(ctx context.Context, tridentProvisioner *v1.TridentProvisioner, opts metav1.CreateOptions) (*v1.TridentProvisioner, error)
	Update(ctx context.Context, tridentProvisioner *v1.TridentProvisioner, opts metav1.UpdateOptions) (*v1.TridentProvisioner, error)
	UpdateStatus(ctx context.Context, tridentProvisioner *v1.TridentProvisioner, opts metav1.UpdateOptions) (*v1.TridentProvisioner, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TridentProvisioner, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TridentProvisionerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentProvisioner, err error)
	TridentProvisionerExpansion
}

// tridentProvisioners implements TridentProvisionerInterface
type tridentProvisioners struct {
	client rest.Interface
	ns     string
}

// newTridentProvisioners returns a TridentProvisioners
func newTridentProvisioners(c *TridentV1Client, namespace string) *tridentProvisioners {
	return &tridentProvisioners{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tridentProvisioner, and returns the corresponding tridentProvisioner object, and an error if there is any.
func (c *tridentProvisioners) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TridentProvisioner, err error) {
	result = &v1.TridentProvisioner{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentprovisioners").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TridentProvisioners that match those selectors.
func (c *tridentProvisioners) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TridentProvisionerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TridentProvisionerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentprovisioners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tridentProvisioners.
func (c *tridentProvisioners) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tridentprovisioners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tridentProvisioner and creates it.  Returns the server's representation of the tridentProvisioner, and an error, if there is any.
func (c *tridentProvisioners) Create(ctx context.Context, tridentProvisioner *v1.TridentProvisioner, opts metav1.CreateOptions) (result *v1.TridentProvisioner, err error) {
	result = &v1.TridentProvisioner{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tridentprovisioners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentProvisioner).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tridentProvisioner and updates it. Returns the server's representation of the tridentProvisioner, and an error, if there is any.
func (c *tridentProvisioners) Update(ctx context.Context, tridentProvisioner *v1.TridentProvisioner, opts metav1.UpdateOptions) (result *v1.TridentProvisioner, err error) {
	result = &v1.TridentProvisioner{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tridentprovisioners").
		Name(tridentProvisioner.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentProvisioner).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tridentProvisioners) UpdateStatus(ctx context.Context, tridentProvisioner *v1.TridentProvisioner, opts metav1.UpdateOptions) (result *v1.TridentProvisioner, err error) {
	result = &v1.TridentProvisioner{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tridentprovisioners").
		Name(tridentProvisioner.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentProvisioner).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tridentProvisioner and deletes it. Returns an error if one occurs.
func (c *tridentProvisioners) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentprovisioners").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tridentProvisioners) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentprovisioners").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tridentProvisioner.
func (c *tridentProvisioners) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentProvisioner, err error) {
	result = &v1.TridentProvisioner{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tridentprovisioners").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
