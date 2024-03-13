package integration

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/aliml92/go-typesense/typesense"
	"github.com/stretchr/testify/require"
)

func TestDocuments_Search(t *testing.T) {
	var (
		err            error
		collectionName = "articles"
	)

	ctx := context.Background()

	// Create `articles` collection
	createArticlesCollection(ctx, t)

	// Import documents
	importArticlesDocuments(ctx, t)

	params := &typesense.SearchParameters{
		Q:       "kubernetes",
		QueryBy: "title,content",
	}

	res, err := client.Documents.Search(ctx, collectionName, params)
	require.NoError(t, err)

	require.Equal(t, 2, *res.Found)
}

func createArticlesCollection(ctx context.Context, t *testing.T) {
	createSchema := &typesense.CollectionSchema{
		Name: "articles",
		Fields: []*typesense.Field{
			{
				Name: "id",
				Type: "string",
			},
			{
				Name: "slug",
				Type: "string",
			},
			{
				Name: "title",
				Type: "string",
			},
			{
				Name: "author",
				Type: "string",
			},
			{
				Name: "content",
				Type: "string",
			},
		},
	}

	_, err := client.Collections.Create(ctx, createSchema)
	require.NoError(t, err)
}

func importArticlesDocuments(ctx context.Context, t *testing.T) {
	var err error
	raw, err := os.ReadFile("testdata/articles.jsonl")
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

	res, err := client.Documents.Import(ctx, "articles", documents, nil)
	require.NoError(t, err)
	for _, r := range res {
		require.True(t, r.Success)
	}
}
