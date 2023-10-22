package typesense

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	baseClient, _ := NewClient(nil, server.URL)

	client := baseClient.WithAPIKey("xyz")

	teardown := func() {
		server.Close()
	}

	return client, mux, teardown
}

func TestNewClient_DefaultConfig(t *testing.T) {
	c, _ := NewClient(nil, "")

	assert.NotNil(t, c)
	assert.NotNil(t, c.client)
	assert.Equal(t, defaultServerURL, c.ServerURL.String())
}

func TestNewClient_CustomConfig(t *testing.T) {
	httpClient := &http.Client{}
	serverURL := "http://custom"

	c, _ := NewClient(httpClient, serverURL)

	assert.Equal(t, httpClient, c.client)
	assert.Equal(t, serverURL, c.ServerURL.String())
}
