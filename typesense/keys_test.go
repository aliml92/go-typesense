package typesense

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeysService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/keys", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{ 
				"keys": [
					{
						"actions": [
							"*"
						],
						"collections": [
							"*"
						],
						"description": "Admin key.",
						"expires_at": 64723363199,
						"id": 0,
						"value_prefix": "NT5F"
					}
				]
			}`)
	})

	ctx := context.Background()
	got, err := client.Keys.List(ctx)
	assert.NoError(t, err)

	want := &ApiKeysResponse{
		Keys: []*ApiKey{
			{
				Actions:     []string{"*"},
				Collections: []string{"*"},
				Description: "Admin key.",
				ExpiresAt:   Int64(64723363199),
				Id:          Int64(0),
				ValuePrefix: String("NT5F"),
			},
		},
	}
	assert.Equal(t, want, got)
}

func TestKeysService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/keys", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"actions": [
					"*"
				],
				"collections": [
					"*"
				],
				"description": "Admin key.",
				"expires_at": 64723363199,
				"id": 0,
				"value": "NT5FrRKmzaItViY67oXRBOJDXLlQ3v7C"
			}`)
	})

	ctx := context.Background()
	body := &ApiKeySchema{
		Actions:     []string{"*"},
		Collections: []string{"*"},
		Description: String("Admin key."),
		ExpiresAt:   Int64(64723363199),
		Value:       String("NT5FrRKmzaItViY67oXRBOJDXLlQ3v7C"),
	}
	got, err := client.Keys.Create(ctx, body)
	assert.NoError(t, err)

	want := &ApiKey{
		Actions:     []string{"*"},
		Collections: []string{"*"},
		Description: "Admin key.",
		ExpiresAt:   Int64(64723363199),
		Id:          Int64(0),
		Value:       String("NT5FrRKmzaItViY67oXRBOJDXLlQ3v7C"),
	}
	assert.Equal(t, want, got)
}

func TestKeysService_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	keyId := 0
	u := fmt.Sprintf("/keys/%d", keyId)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"actions": [
					"*"
				],
				"collections": [
					"*"
				],
				"description": "Admin key.",
				"expires_at": 64723363199,
				"id": 0,
				"value_prefix": "NT5F"
			}`)
	})

	ctx := context.Background()
	got, err := client.Keys.Get(ctx, keyId)
	assert.NoError(t, err)

	want := &ApiKey{
		Actions:     []string{"*"},
		Collections: []string{"*"},
		Description: "Admin key.",
		ExpiresAt:   Int64(64723363199),
		Id:          Int64(0),
		ValuePrefix: String("NT5F"),
	}
	assert.Equal(t, want, got)
}

func TestKeysService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	keyId := 0
	u := fmt.Sprintf("/keys/%d", keyId)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"id": 0
			}`)
	})

	ctx := context.Background()
	got, err := client.Keys.Delete(ctx, keyId)
	assert.NoError(t, err)

	want := &ApiKey{
		Id: Int64(0),
	}
	assert.Equal(t, want, got)
}
