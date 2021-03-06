// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	apps_v1 "github.com/kubernetes-sigs/kubebuilder/test/projects/validations/pkg/apis/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTestKinds implements TestKindInterface
type FakeTestKinds struct {
	Fake *FakeAppsV1
	ns   string
}

var testkindsResource = schema.GroupVersionResource{Group: "apps.validation.com", Version: "v1", Resource: "testkinds"}

var testkindsKind = schema.GroupVersionKind{Group: "apps.validation.com", Version: "v1", Kind: "TestKind"}

// Get takes name of the testKind, and returns the corresponding testKind object, and an error if there is any.
func (c *FakeTestKinds) Get(name string, options v1.GetOptions) (result *apps_v1.TestKind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(testkindsResource, c.ns, name), &apps_v1.TestKind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.TestKind), err
}

// List takes label and field selectors, and returns the list of TestKinds that match those selectors.
func (c *FakeTestKinds) List(opts v1.ListOptions) (result *apps_v1.TestKindList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(testkindsResource, testkindsKind, c.ns, opts), &apps_v1.TestKindList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apps_v1.TestKindList{}
	for _, item := range obj.(*apps_v1.TestKindList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested testKinds.
func (c *FakeTestKinds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(testkindsResource, c.ns, opts))

}

// Create takes the representation of a testKind and creates it.  Returns the server's representation of the testKind, and an error, if there is any.
func (c *FakeTestKinds) Create(testKind *apps_v1.TestKind) (result *apps_v1.TestKind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(testkindsResource, c.ns, testKind), &apps_v1.TestKind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.TestKind), err
}

// Update takes the representation of a testKind and updates it. Returns the server's representation of the testKind, and an error, if there is any.
func (c *FakeTestKinds) Update(testKind *apps_v1.TestKind) (result *apps_v1.TestKind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(testkindsResource, c.ns, testKind), &apps_v1.TestKind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.TestKind), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTestKinds) UpdateStatus(testKind *apps_v1.TestKind) (*apps_v1.TestKind, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(testkindsResource, "status", c.ns, testKind), &apps_v1.TestKind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.TestKind), err
}

// Delete takes name of the testKind and deletes it. Returns an error if one occurs.
func (c *FakeTestKinds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(testkindsResource, c.ns, name), &apps_v1.TestKind{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTestKinds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(testkindsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &apps_v1.TestKindList{})
	return err
}

// Patch applies the patch and returns the patched testKind.
func (c *FakeTestKinds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps_v1.TestKind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(testkindsResource, c.ns, name, data, subresources...), &apps_v1.TestKind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.TestKind), err
}
