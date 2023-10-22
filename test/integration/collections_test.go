package integration

import (
	"context"
	"testing"

	"github.com/aliml92/typesense/typesense"
	"github.com/stretchr/testify/require"
)

func TestCollections_CRUD(t *testing.T) {
	var err error

	collectionName := "companies"
	ctx := context.Background()

	// Create collection
	createSchema := &typesense.CollectionSchema{
		Name: collectionName,
		Fields: []*typesense.Field{
			{
				Name: "company_name",
				Type: "string",
			},
			{
				Name: "num_employees",
				Type: "int32",
			},
			{
				Name:  "country",
				Type:  "string",
				Facet: typesense.Bool(true),
			},
		},
		DefaultSortingField: typesense.String("num_employees"),
	}

	_, err = client.Collections.Create(ctx, createSchema)
	require.NoError(t, err)

	// List all collections
	list, err := client.Collections.List(ctx)
	require.NoError(t, err)
	require.Equal(t, 1, len(list))

	// Get a collection
	collection, err := client.Collections.Get(ctx, collectionName)
	require.NoError(t, err)
	require.Equal(t, createSchema.Name, collection.Name)
	require.Equal(t, createSchema.DefaultSortingField, collection.DefaultSortingField)
	require.Equal(t, int64(0), *collection.NumDocuments)

	// Update collection
	updateSchema := &typesense.CollectionUpdateSchema{
		Fields: []*typesense.Field{
			{
				Name: "country",
				Drop: typesense.Bool(true),
			},
			{
				Name:  "location",
				Type:  "string",
				Facet: typesense.Bool(true),
			},
		},
	}
	_, err = client.Collections.Update(ctx, collectionName, updateSchema)
	require.NoError(t, err)

	// Delete collection
	_, err = client.Collections.Delete(ctx, collectionName)
	require.NoError(t, err)
}
