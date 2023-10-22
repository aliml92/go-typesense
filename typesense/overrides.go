package typesense

import (
	"context"
	"fmt"
)

type OverridesService service

func (s *OverridesService) List(ctx context.Context, collectionName string) (*SearchOverridesResponse, error) {
	u := fmt.Sprintf("/collections/%s/overrides", collectionName)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &SearchOverridesResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *OverridesService) Get(ctx context.Context, collectionName, overrideId string) (*SearchOverride, error) {
	u := fmt.Sprintf("/collections/%s/overrides/%s", collectionName, overrideId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &SearchOverride{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *OverridesService) Upsert(ctx context.Context, collectionName, overrideId string, body *SearchOverrideSchema) (*SearchOverride, error) {
	u := fmt.Sprintf("/collections/%s/overrides/%s", collectionName, overrideId)
	req, err := s.client.NewRequest("PUT", u, body)
	if err != nil {
		return nil, err
	}

	res := &SearchOverride{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type DeleteOverrideResponse struct {
	ID string `json:"id"`
}

func (s *OverridesService) Delete(ctx context.Context, collectionName, overrideId string) (*DeleteOverrideResponse, error) {
	u := fmt.Sprintf("/collections/%s/overrides/%s", collectionName, overrideId)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	res := &DeleteOverrideResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
