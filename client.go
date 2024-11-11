package evolution

import (
	"net/http"
)

type Client struct {
	apiKey     string
	baseUrl    string
	httpClient *http.Client
}

// Option is a function that configures a client
type Option func(*Client)

// NewClient creates a new client with the provided options
func NewClient(opts ...Option) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithBaseUrl sets the base URL of the client
func WithBaseUrl(baseUrl string) Option {
	return func(c *Client) {
		c.baseUrl = baseUrl
	}
}

// WithApiKey sets the api key id of the client
func WithApiKey(sessionKey string) Option {
	return func(c *Client) {
		c.apiKey = sessionKey
	}
}

// WithHttpClient sets the http client of the client
func WithHttpClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
