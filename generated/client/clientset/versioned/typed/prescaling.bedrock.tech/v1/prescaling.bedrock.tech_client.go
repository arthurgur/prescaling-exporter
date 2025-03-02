// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"net/http"

	"github.com/arthurgur/prescaling-exporter/generated/client/clientset/versioned/scheme"
	v1 "github.com/arthurgur/prescaling-exporter/pkg/apis/prescaling.bedrock.tech/v1"
	rest "k8s.io/client-go/rest"
)

type PrescalingV1Interface interface {
	RESTClient() rest.Interface
	PrescalingEventsGetter
}

// PrescalingV1Client is used to interact with features provided by the prescaling.bedrock.tech group.
type PrescalingV1Client struct {
	restClient rest.Interface
}

func (c *PrescalingV1Client) PrescalingEvents(namespace string) PrescalingEventInterface {
	return newPrescalingEvents(c, namespace)
}

// NewForConfig creates a new PrescalingV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*PrescalingV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new PrescalingV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*PrescalingV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &PrescalingV1Client{client}, nil
}

// NewForConfigOrDie creates a new PrescalingV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *PrescalingV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new PrescalingV1Client for the given RESTClient.
func New(c rest.Interface) *PrescalingV1Client {
	return &PrescalingV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *PrescalingV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
