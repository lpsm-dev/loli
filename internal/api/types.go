package api

import "net/http"

// Client struct - is an http client that is configured for API.
type Client struct {
	*http.Client        // HTTP Client pointer
	BaseURL      string // HTTP Request URL
	ContentType  string // HTTP Request Content Type
}
