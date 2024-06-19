package typesense

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const apiKey = "xyz"

func setup() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client, _ := NewClient(nil, server.URL, apiKey)

	teardown := func() {
		server.Close()
	}

	return client, mux, teardown
}

func TestNewClient_DefaultConfig(t *testing.T) {
	c, _ := NewClient(nil, "", apiKey)

	assert.NotNil(t, c)
	assert.NotNil(t, c.client)
	assert.Equal(t, defaultServerURL, c.serverURL.String(), apiKey)
}

func TestNewClient_CustomConfig(t *testing.T) {
	httpClient := &http.Client{}
	serverURL := "http://custom"

	c, _ := NewClient(httpClient, serverURL, apiKey)

	assert.Equal(t, httpClient, c.client)
	assert.Equal(t, serverURL, c.serverURL.String())
}
