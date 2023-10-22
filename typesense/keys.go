package typesense

import (
	"context"
	"fmt"
)

type KeysService service

func (s *KeysService) List(ctx context.Context) (*ApiKeysResponse, error) {
	u := "/keys"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &ApiKeysResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *KeysService) Create(ctx context.Context, body *ApiKeySchema) (*ApiKey, error) {
	u := "/keys"
	req, err := s.client.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}

	res := &ApiKey{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *KeysService) Get(ctx context.Context, keyId int) (*ApiKey, error) {
	u := fmt.Sprintf("/keys/%d", keyId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &ApiKey{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *KeysService) Delete(ctx context.Context, keyId int) (*ApiKey, error) {
	u := fmt.Sprintf("/keys/%d", keyId)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	res := &ApiKey{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
