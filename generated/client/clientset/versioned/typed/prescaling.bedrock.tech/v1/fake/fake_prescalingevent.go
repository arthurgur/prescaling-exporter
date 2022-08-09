// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	prescalingbedrocktechv1 "github.com/BedrockStreaming/prescaling-exporter/pkg/apis/prescaling.bedrock.tech/v1"
)

// FakePrescalingEvents implements PrescalingEventInterface
type FakePrescalingEvents struct {
	Fake *FakePrescalingV1
	ns   string
}

var prescalingeventsResource = schema.GroupVersionResource{Group: "prescaling.bedrock.tech", Version: "v1", Resource: "prescalingevents"}

var prescalingeventsKind = schema.GroupVersionKind{Group: "prescaling.bedrock.tech", Version: "v1", Kind: "PrescalingEvent"}

// Get takes name of the prescalingEvent, and returns the corresponding prescalingEvent object, and an error if there is any.
func (c *FakePrescalingEvents) Get(ctx context.Context, name string, options v1.GetOptions) (result *prescalingbedrocktechv1.PrescalingEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(prescalingeventsResource, c.ns, name), &prescalingbedrocktechv1.PrescalingEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*prescalingbedrocktechv1.PrescalingEvent), err
}

// List takes label and field selectors, and returns the list of PrescalingEvents that match those selectors.
func (c *FakePrescalingEvents) List(ctx context.Context, opts v1.ListOptions) (result *prescalingbedrocktechv1.PrescalingEventList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(prescalingeventsResource, prescalingeventsKind, c.ns, opts), &prescalingbedrocktechv1.PrescalingEventList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &prescalingbedrocktechv1.PrescalingEventList{ListMeta: obj.(*prescalingbedrocktechv1.PrescalingEventList).ListMeta}
	for _, item := range obj.(*prescalingbedrocktechv1.PrescalingEventList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested prescalingEvents.
func (c *FakePrescalingEvents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(prescalingeventsResource, c.ns, opts))

}

// Create takes the representation of a prescalingEvent and creates it.  Returns the server's representation of the prescalingEvent, and an error, if there is any.
func (c *FakePrescalingEvents) Create(ctx context.Context, prescalingEvent *prescalingbedrocktechv1.PrescalingEvent, opts v1.CreateOptions) (result *prescalingbedrocktechv1.PrescalingEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(prescalingeventsResource, c.ns, prescalingEvent), &prescalingbedrocktechv1.PrescalingEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*prescalingbedrocktechv1.PrescalingEvent), err
}

// Update takes the representation of a prescalingEvent and updates it. Returns the server's representation of the prescalingEvent, and an error, if there is any.
func (c *FakePrescalingEvents) Update(ctx context.Context, prescalingEvent *prescalingbedrocktechv1.PrescalingEvent, opts v1.UpdateOptions) (result *prescalingbedrocktechv1.PrescalingEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(prescalingeventsResource, c.ns, prescalingEvent), &prescalingbedrocktechv1.PrescalingEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*prescalingbedrocktechv1.PrescalingEvent), err
}

// Delete takes name of the prescalingEvent and deletes it. Returns an error if one occurs.
func (c *FakePrescalingEvents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(prescalingeventsResource, c.ns, name), &prescalingbedrocktechv1.PrescalingEvent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrescalingEvents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(prescalingeventsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &prescalingbedrocktechv1.PrescalingEventList{})
	return err
}

// Patch applies the patch and returns the patched prescalingEvent.
func (c *FakePrescalingEvents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *prescalingbedrocktechv1.PrescalingEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(prescalingeventsResource, c.ns, name, pt, data, subresources...), &prescalingbedrocktechv1.PrescalingEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*prescalingbedrocktechv1.PrescalingEvent), err
}
