package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequestSetsDefaultHeaders(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `ok`)
	}))
	defer ts.Close()

	UserAgent = "BogusAgent"

	testCases := []struct {
		desc        string
		client      *Client
		contentType string
	}{
		{
			desc:        "User defaults",
			client:      &Client{},
			contentType: "application/json",
		},
		{
			desc: "Override defaults",
			client: &Client{
				BaseURL:     "http://example.com",
				ContentType: "bogus",
			},
			contentType: "bogus",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			req, err := tc.client.NewRequest("GET", ts.URL, nil)
			assert.NoError(t, err)
			assert.Equal(t, "BogusAgent", req.Header.Get("User-Agent"))
			assert.Equal(t, tc.contentType, req.Header.Get("Content-Type"))
		})
	}
}

func TestDo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)

		fmt.Fprint(w, `{"hello": "world"}`)
	}))
	defer ts.Close()

	type payload struct {
		Hello string `json:"hello"`
	}

	client := &Client{}

	req, err := client.NewRequest("GET", ts.URL, nil)
	assert.NoError(t, err)

	res, err := client.Do(req)
	assert.NoError(t, err)

	var body payload
	err = json.NewDecoder(res.Body).Decode(&body)
	assert.NoError(t, err)

	assert.Equal(t, "world", body.Hello)
}
