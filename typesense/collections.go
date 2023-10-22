package typesense

import (
	"context"
	"fmt"
)

type CollectionsService service

func (s *CollectionsService) List(ctx context.Context) ([]*CollectionSchema, error) {
	req, err := s.client.NewRequest("GET", "/collections", nil)
	if err != nil {
		return nil, err
	}

	var res []*CollectionSchema
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CollectionsService) Create(ctx context.Context, body *CollectionSchema) (*Collection, error) {
	req, err := s.client.NewRequest("POST", "/collections", body)
	if err != nil {
		return nil, err
	}

	res := &Collection{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CollectionsService) Get(ctx context.Context, collectionName string) (*Collection, error) {
	u := fmt.Sprintf("/collections/%s", collectionName)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &Collection{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CollectionsService) Update(ctx context.Context, collectionName string, body *CollectionUpdateSchema) (*CollectionUpdateSchema, error) {
	u := fmt.Sprintf("/collections/%s", collectionName)
	req, err := s.client.NewRequest("PATCH", u, body)
	if err != nil {
		return nil, err
	}

	res := &CollectionUpdateSchema{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CollectionsService) Delete(ctx context.Context, collectionName string) (*Collection, error) {
	u := fmt.Sprintf("/collections/%s", collectionName)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	res := &Collection{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
