package integration

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aliml92/typesense/typesense"
	"github.com/stretchr/testify/require"
)

func TestDocuments_MultiSearch(t *testing.T) {
	var err error

	collectionName1 := "clothes"
	collectionName2 := "shoes"
	collectionName3 := "bags" // bags do no exist

	ctx := context.Background()

	// Create collections
	createCollection(ctx, t, collectionName1)
	createCollection(ctx, t, collectionName2)

	// Import documents
	importDocuments(ctx, t, collectionName1, "testdata/clothes.jsonl")
	importDocuments(ctx, t, collectionName2, "testdata/shoes.jsonl")

	time.Sleep(time.Second) // wait until documents get indexed

	// search for `clothes` and `shoes` collections
	body := &typesense.MultiSearchSearchesParameter{
		Searches: []typesense.MultiSearchCollectionParameters{
			{
				Collection: collectionName1,
				Q:          typesense.String("polo"),
				QueryBy:    typesense.String("name,description"),
			},
			{
				Collection: collectionName2,
				Q:          typesense.String("sneakers"),
				QueryBy:    typesense.String("name,description"),
			},
		},
	}

	params := &typesense.MultiSearchParameters{
		Page:    typesense.Int(0),
		PerPage: typesense.Int(10),
	}

	res, err := client.Documents.MultiSearch(ctx, body, params)
	require.NoError(t, err)

	require.Nil(t, res.Results[0].SearchError)
	require.Equal(t, 1, *res.Results[0].SearchResult.Found)

	require.Nil(t, res.Results[1].SearchError)
	require.Equal(t, 2, *res.Results[1].SearchResult.Found)

	// search for `clothes` and `bags` collection
	// since `bags` collection do not exist, `multi_search` response will
	// also contain an error response
	body = &typesense.MultiSearchSearchesParameter{
		Searches: []typesense.MultiSearchCollectionParameters{
			{
				Collection: collectionName1,
				Q:          typesense.String("polo"),
				QueryBy:    typesense.String("name,description"),
			},
			{
				Collection: collectionName3,
				Q:          typesense.String("backpack"),
				QueryBy:    typesense.String("name, description"),
			},
		},
	}

	params = &typesense.MultiSearchParameters{
		Page:    typesense.Int(0),
		PerPage: typesense.Int(10),
	}

	res, err = client.Documents.MultiSearch(ctx, body, params)
	require.NoError(t, err)

	require.Nil(t, res.Results[0].SearchError)
	require.Equal(t, 1, *res.Results[0].SearchResult.Found)

	require.Nil(t, res.Results[1].SearchResult)
	require.Equal(t, 404, res.Results[1].SearchError.Code)
	require.Equal(t, "Not found.", res.Results[1].SearchError.Error)
}

func createCollection(ctx context.Context, t *testing.T, cn string) {
	createSchema := &typesense.CollectionSchema{
		Name: cn,
		Fields: []*typesense.Field{
			{
				Name: "id",
				Type: "string",
			},
			{
				Name: "name",
				Type: "string",
			},
			{
				Name:  "category",
				Type:  "string",
				Facet: typesense.Bool(true),
			},
			{
				Name: "color",
				Type: "string",
			},
			{
				Name: "size",
				Type: "string",
			},
			{
				Name: "price",
				Type: "float",
			},

			{
				Name: "description",
				Type: "string",
			},
		},
		DefaultSortingField: typesense.String("price"),
	}

	_, err := client.Collections.Create(ctx, createSchema)
	require.NoError(t, err)
}

func importDocuments(ctx context.Context, t *testing.T, cn, importFile string) {
	var err error
	raw, err := os.ReadFile(importFile)
	if err != nil {
		require.FailNow(t, err.Error())
	}

	lines := strings.Split(string(raw), "\n")
	var documents []map[string]interface{}

	for _, line := range lines {
		var data map[string]interface{}
		if err = json.Unmarshal([]byte(line), &data); err != nil {
			require.FailNow(t, err.Error())
		}

		documents = append(documents, data)
	}

	res, err := client.Documents.Import(ctx, cn, documents, nil)
	require.NoError(t, err)
	for _, r := range res {
		require.True(t, r.Success)
	}
}
