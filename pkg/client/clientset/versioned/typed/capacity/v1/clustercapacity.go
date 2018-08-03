// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1 "github.com/supergiant/capacity/pkg/apis/capacity/v1"
	scheme "github.com/supergiant/capacity/pkg/client/clientset/versioned/scheme"
)

// ClusterCapacitiesGetter has a method to return a ClusterCapacityInterface.
// A group's client should implement this interface.
type ClusterCapacitiesGetter interface {
	ClusterCapacities(namespace string) ClusterCapacityInterface
}

// ClusterCapacityInterface has methods to work with ClusterCapacity resources.
type ClusterCapacityInterface interface {
	Create(*v1.ClusterCapacity) (*v1.ClusterCapacity, error)
	Update(*v1.ClusterCapacity) (*v1.ClusterCapacity, error)
	UpdateStatus(*v1.ClusterCapacity) (*v1.ClusterCapacity, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.ClusterCapacity, error)
	List(opts meta_v1.ListOptions) (*v1.ClusterCapacityList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterCapacity, err error)
	ClusterCapacityExpansion
}

// clusterCapacities implements ClusterCapacityInterface
type clusterCapacities struct {
	client rest.Interface
	ns     string
}

// newClusterCapacities returns a ClusterCapacities
func newClusterCapacities(c *CapacityV1Client, namespace string) *clusterCapacities {
	return &clusterCapacities{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterCapacity, and returns the corresponding clusterCapacity object, and an error if there is any.
func (c *clusterCapacities) Get(name string, options meta_v1.GetOptions) (result *v1.ClusterCapacity, err error) {
	result = &v1.ClusterCapacity{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustercapacities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterCapacities that match those selectors.
func (c *clusterCapacities) List(opts meta_v1.ListOptions) (result *v1.ClusterCapacityList, err error) {
	result = &v1.ClusterCapacityList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustercapacities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterCapacities.
func (c *clusterCapacities) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clustercapacities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a clusterCapacity and creates it.  Returns the server's representation of the clusterCapacity, and an error, if there is any.
func (c *clusterCapacities) Create(clusterCapacity *v1.ClusterCapacity) (result *v1.ClusterCapacity, err error) {
	result = &v1.ClusterCapacity{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clustercapacities").
		Body(clusterCapacity).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterCapacity and updates it. Returns the server's representation of the clusterCapacity, and an error, if there is any.
func (c *clusterCapacities) Update(clusterCapacity *v1.ClusterCapacity) (result *v1.ClusterCapacity, err error) {
	result = &v1.ClusterCapacity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustercapacities").
		Name(clusterCapacity.Name).
		Body(clusterCapacity).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterCapacities) UpdateStatus(clusterCapacity *v1.ClusterCapacity) (result *v1.ClusterCapacity, err error) {
	result = &v1.ClusterCapacity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustercapacities").
		Name(clusterCapacity.Name).
		SubResource("status").
		Body(clusterCapacity).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterCapacity and deletes it. Returns an error if one occurs.
func (c *clusterCapacities) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustercapacities").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterCapacities) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustercapacities").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterCapacity.
func (c *clusterCapacities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterCapacity, err error) {
	result = &v1.ClusterCapacity{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clustercapacities").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}