package typesense

import (
	"context"
)

type AnalyticsEventsService service

type AnalyticsEvent struct {
	Type string `json:"type"`
	Data struct {
		Q           string   `json:"q"`
		Collections []string `json:"collections"`
	} `json:"data"`
}

type AnalyticsEventCreateResponse struct {
	OK bool `json:"ok"`
}

func (s *AnalyticsEventsService) Create(ctx context.Context, body *AnalyticsEvent) (*AnalyticsEventCreateResponse, error) {
	u := "/analytics/events"
	req, err := s.client.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}

	res := &AnalyticsEventCreateResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
