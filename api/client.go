package api

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/lpmatos/loli/debug"
)

var (
	// UserAgent variable - lets the API know where the call is being made from.
	// For more information: https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Headers/User-Agent
	UserAgent = "github.com/lpmatos/loli"

	// TimeoutInSeconds variable - is the timeout the default HTTP client will use.
	// For more information: https://stackoverflow.com/questions/16895294/how-to-set-timeout-for-http-get-requests-in-golang
	TimeoutInSeconds = 60

	// InsecureSkipVerify controls whether a client verifies the server's certificate chain and host name.
	// For more information: https://golang.org/pkg/crypto/tls/
	InsecureSkipVerify = false

	// HTTPClient variable - is the client used to make HTTP calls in loli cli.
	// For more information: https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
	HTTPClient = &http.Client{
		Timeout: time.Duration(TimeoutInSeconds) * time.Second,
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
				InsecureSkipVerify: InsecureSkipVerify,
			},
		},
	}
)

// NewClient returns a trace.moe API client.
func NewClient(baseURL string) (*Client, error) {
	return &Client{
		Client:  HTTPClient,
		BaseURL: baseURL,
	}, nil
}

// NewRequest returns an http.Request with information for the trace.moe API.
// For more information: https://golang.org/pkg/net/http/
func (c *Client) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	if c.Client == nil {
		c.Client = HTTPClient
	}

	request, error := http.NewRequest(method, url, body)
	if error != nil {
		return nil, error
	}

	request.Header.Set("User-Agent", UserAgent)

	if c.ContentType == "" {
		request.Header.Set("Content-Type", "application/json")
	} else {
		request.Header.Set("Content-Type", c.ContentType)
	}

	return request, nil
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
