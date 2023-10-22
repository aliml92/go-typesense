package typesense

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyticsEventsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	u := "/analytics/events"
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"ok": true
		}`)
	})

	ctx := context.Background()
	body := &AnalyticsEvent{
		Type: "search",
		Data: struct {
			Q           string   `json:"q"`
			Collections []string `json:"collections"`
		}{
			Q:           "Nike shoes",
			Collections: []string{"products"},
		},
	}
	got, err := client.AnalyticsEvents.Create(ctx, body)
	assert.NoError(t, err)

	want := &AnalyticsEventCreateResponse{
		OK: true,
	}
	assert.Equal(t, want, got)
}
