package api

import (
	"io"
	"net/http"
	"time"
)

var (
	// UserAgent lets the API know where the call is being made from.
	// It's overridden from the root command so that we can set the version.
	UserAgent = "github.com/lpmatos/loli"

	// TimeoutInSeconds variable - is the timeout the default HTTP client will use.
	TimeoutInSeconds = 60

	// HTTPClient variable - is the client used to make HTTP calls in loli cli.
	HTTPClient = &http.Client{
		Timeout: time.Duration(TimeoutInSeconds) * time.Second,
	}
)

// Client struct - is an http client that is configured for trace.
type Client struct {
	*http.Client
	BaseURL     string
	ContentType string
}

// NewClient returns an trace API client.
func NewClient(baseURL string) (*Client, error) {
	return &Client{
		Client:  HTTPClient,
		BaseURL: baseURL,
	}, nil
}

// NewRequest returns an http.Request with information for the trace API.
func (c *Client) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	if c.Client == nil {
		c.Client = HTTPClient
	}

	request, error := http.NewRequest(method, url, body)
	if error != nil {
		return nil, error
	}

	request.Header.Set("Content-Type", UserAgent)

	if c.ContentType == "" {
		request.Header.Set("Content-Type", "application/json")
	} else {
		request.Header.Set("Content-Type", c.ContentType)
	}

	return request, nil
}
