package typesense

import (
	"context"
	"fmt"
)

type PresetsService service

type Preset struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type PresetUpsertSchema struct {
	Value interface{} `json:"value"`
}

type PresetListResponse struct {
	Presets []*Preset `json:"presets"`
}

func (s *PresetsService) List(ctx context.Context) (*PresetListResponse, error) {
	u := "/presets"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &PresetListResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PresetsService) Upsert(ctx context.Context, presetName string, body *PresetUpsertSchema) (*Preset, error) {
	u := fmt.Sprintf("/presets/%s", presetName)
	req, err := s.client.NewRequest("PUT", u, body)
	if err != nil {
		return nil, err
	}

	res := &Preset{}
	switch body.Value.(type) {
	case *SearchParameters:
		res.Value = &SearchParameters{}
	case *MultiSearchSearchesParameter:
		res.Value = &MultiSearchSearchesParameter{}
	default:
		return nil, fmt.Errorf("invalid type for body.Value")
	}

	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PresetsService) Get(ctx context.Context, presetName string) (*Preset, error) {
	u := fmt.Sprintf("/presets/%s", presetName)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &Preset{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PresetsService) Delete(ctx context.Context, presetName string) (*Preset, error) {
	u := fmt.Sprintf("/presets/%s", presetName)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	res := &Preset{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
