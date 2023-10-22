package typesense

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPresetService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/presets", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
		{
			"presets": [
				{
					"name": "listing_view",
					"value": {
						"searches": [
							{
								"collection": "products",
								"q": "*",
								"sort_by": "popularity"
							}
						]
					}
				}
			]
		}`)
	})

	ctx := context.Background()
	got, err := client.Presets.List(ctx)
	assert.NoError(t, err)

	want := &PresetListResponse{
		Presets: []*Preset{
			{
				Name: "listing_view",
				Value: &MultiSearchSearchesParameter{
					Searches: []MultiSearchCollectionParameters{
						{
							Collection: "products",
							Q:          String("*"),
							SortBy:     String("popularity"),
						},
					},
				},
			},
		},
	}

	assert.Equal(t, want.Presets[0].Name, got.Presets[0].Name)
}

func TestPresetService_Upsert(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	presetName := "listing_view"
	u := fmt.Sprintf("/presets/%s", presetName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{ 
				"name": "listing_view",
				"value": {
					"searches": [
						{
							"collection": "products",
							"q": "*",
							"sort_by": "popularity"
						}
					]
				}
			}`)
	})

	ctx := context.Background()
	body := &PresetUpsertSchema{
		Value: &MultiSearchSearchesParameter{
			Searches: []MultiSearchCollectionParameters{
				{
					Collection: "products",
					Q:          String("*"),
					SortBy:     String("popularity"),
				},
			},
		},
	}
	got, err := client.Presets.Upsert(ctx, presetName, body)
	assert.NoError(t, err)

	want := &Preset{
		Name: presetName,
		Value: &MultiSearchSearchesParameter{
			Searches: []MultiSearchCollectionParameters{
				{
					Collection: "products",
					Q:          String("*"),
					SortBy:     String("popularity"),
				},
			},
		},
	}

	assert.Equal(t, want, got)
}
