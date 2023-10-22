package typesense

import (
	"context"
	"fmt"
)

type AliasesService service

func (s *AliasesService) List(ctx context.Context) ([]*CollectionAlias, error) {
	u := "/aliases"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res []*CollectionAlias
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AliasesService) Upsert(ctx context.Context, aliasName string, body *CollectionAliasSchema) (*CollectionAlias, error) {
	u := fmt.Sprintf("/aliases/%s", aliasName)
	req, err := s.client.NewRequest("PUT", u, body)
	if err != nil {
		return nil, err
	}

	res := &CollectionAlias{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AliasesService) Get(ctx context.Context, aliasName string) (*CollectionAlias, error) {
	u := fmt.Sprintf("/aliases/%s", aliasName)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &CollectionAlias{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AliasesService) Delete(ctx context.Context, aliasName string) (*CollectionAlias, error) {
	u := fmt.Sprintf("/aliases/%s", aliasName)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	res := &CollectionAlias{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
