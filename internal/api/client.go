package api

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/lpmatos/loli/internal/constants"
	"github.com/lpmatos/loli/internal/debug"
)

var (
	// HTTPClient variable - is the client used to make HTTP calls in loli cli.
	// For more information: https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
	HTTPClient = &http.Client{
		Timeout: time.Duration(constants.TimeoutInSeconds) * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: constants.InsecureSkipVerify,
			},
		},
	}
)

// NewClient returns a API client.
func NewClient(baseURL string) (*Client, error) {
	return &Client{
		Client:  HTTPClient,
		BaseURL: baseURL,
	}, nil
}

// NewRequest returns an http.Request with information for the API. For more information: https://golang.org/pkg/net/http/
func (c *Client) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	if c.Client == nil {
		c.Client = HTTPClient
	}

	req, error := http.NewRequest(method, url, body)
	if error != nil {
		return nil, error
	}

	req.Header.Set("User-Agent", constants.UserAgent)

	if c.ContentType == "" {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", c.ContentType)
	}

	return req, nil
}

// Do performs an http.Request and optionally parses the response body into the given interface.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	debug.DumpRequest(req)

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	debug.DumpResponse(res)
	return res, nil
}
