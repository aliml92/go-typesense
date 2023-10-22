package typesense

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperationsServiceService_Snapshot(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/operations/snapshot", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"success": true
			}
		`)
	})

	ctx := context.Background()
	opts := &TakeSnapshotParams{
		SnapshotPath: "/tmp/typesense-data-snapshot",
	}
	got, err := client.Operations.Snapshot(ctx, opts)
	assert.NoError(t, err)
	want := &SuccessStatus{
		Success: true,
	}

	assert.Equal(t, want, got)
}

func TestOperationsServiceService_Vote(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/operations/vote", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"success": true
			}
		`)
	})

	ctx := context.Background()
	got, err := client.Operations.Vote(ctx)
	assert.NoError(t, err)
	want := &SuccessStatus{
		Success: true,
	}

	assert.Equal(t, want, got)
}

func TestOperationsServiceService_ClearCache(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/operations/cache/clear", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"success": true
			}
		`)
	})

	ctx := context.Background()
	got, err := client.Operations.ClearCache(ctx)
	assert.NoError(t, err)
	want := &SuccessStatus{
		Success: true,
	}

	assert.Equal(t, want, got)
}

func TestOperationsServiceService_CompactDB(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/operations/compact/db", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"success": true
			}
		`)
	})

	ctx := context.Background()
	got, err := client.Operations.CompactDB(ctx)
	assert.NoError(t, err)
	want := &SuccessStatus{
		Success: true,
	}

	assert.Equal(t, want, got)
}

func TestOperationsServiceService_ResetPeers(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/operations/reset_peers", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"success": true
			}
		`)
	})

	ctx := context.Background()
	got, err := client.Operations.ResetPeers(ctx)
	assert.NoError(t, err)
	want := &SuccessStatus{
		Success: true,
	}

	assert.Equal(t, want, got)
}
