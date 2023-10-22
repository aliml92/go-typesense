package integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/aliml92/typesense/typesense"
	"github.com/stretchr/testify/require"
)

func TestRateLimits(t *testing.T) {
	var err error

	ctx := context.Background()

	// Create a key
	keySchema := &typesense.ApiKeySchema{
		Actions:     []string{"*"},
		Collections: []string{"*"},
		Description: typesense.String("a user key"),
		Value:       typesense.String("abc"),
	}

	_, err = client.Keys.Create(ctx, keySchema)
	require.NoError(t, err)

	// Test Create
	ruleSchema := &typesense.RateLimitRuleSchema{
		Action:                 typesense.THROTTLE,
		ApiKeys:                []string{"abc"},
		MaxRequests1m:          typesense.Int(10),
		MaxRequests1h:          typesense.Int(-1),
		AutoBan1mThreshold:     typesense.Int(1),
		AutoBan1mDurationHours: typesense.Int(1),
	}

	createRes, err := client.RateLimits.Create(ctx, ruleSchema)
	require.NoError(t, err)
	require.Equal(t, "Rule added successfully.", createRes.Message)

	// create new client with an api key
	newClient := baseClient.WithAPIKey("abc")

	// Make 10 requests in 1m to trigger auto ban
	for i := 0; i < 10; i++ {
		_, err = newClient.Collections.List(ctx)
		require.NoError(t, err)
	}

	// Now the following request will raise TooManyRequests error
	_, err = newClient.Collections.List(ctx)
	require.Error(t, err)

	// go:nolint
	ae, ok := err.(*typesense.ApiError)
	require.True(t, ok)
	require.Equal(t, http.StatusTooManyRequests, ae.StatusCode)
	require.Equal(t, "Rate limit exceeded or blocked", ae.Body.Message)
}
