package typesense

import (
	"context"
	"fmt"
)

type AnalyticsRuleType string

const (
	POPULAR_QUERIES_TYPE AnalyticsRuleType = "popular_queries"
)

type AnalyticsRulesService service

type AnalyticsRuleDeleteResponse struct {
	Name string `json:"name"`
}

type AnalyticsRuleListResponse struct {
	Rules []*AnalyticsRule `json:"rules"`
}

type AnalyticsRule struct {
	Name   string              `json:"name"`
	Type   AnalyticsRuleType   `json:"type"` // only "popular_queries" is supported
	Params AnalyticsRuleParams `json:"params"`
}

type AnalyticsRuleParams struct {
	Source struct {
		Collections []string `json:"collections"`
	} `json:"source"`
	Destination struct {
		Collection string `json:"collection"`
	} `json:"destination"`
	Limit int `json:"limit"`
}

type AnalyticsRuleUpsertSchema struct {
	Type   AnalyticsRuleType `json:"type"` // only "popular_queries" is supported
	Params struct {
		Source struct {
			Collections []string `json:"collections"`
		} `json:"source"`
		Destination struct {
			Collection string `json:"collection"`
		} `json:"destination"`
		Limit int `json:"limit"`
	} `json:"params"`
}

func (s *AnalyticsRulesService) List(ctx context.Context) (*AnalyticsRuleListResponse, error) {
	u := "/analytics/rules"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &AnalyticsRuleListResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AnalyticsRulesService) Create(ctx context.Context, body *AnalyticsRule) (*AnalyticsRule, error) {
	u := "/analytics/rules"
	req, err := s.client.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}

	res := &AnalyticsRule{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AnalyticsRulesService) Get(ctx context.Context, ruleName string) (*AnalyticsRule, error) {
	u := fmt.Sprintf("/analytics/rules/%s", ruleName)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &AnalyticsRule{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AnalyticsRulesService) Upsert(ctx context.Context, ruleName string, body *AnalyticsRuleUpsertSchema) (*AnalyticsRule, error) {
	u := fmt.Sprintf("/analytics/rules/%s", ruleName)
	req, err := s.client.NewRequest("PUT", u, body)
	if err != nil {
		return nil, err
	}

	res := &AnalyticsRule{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AnalyticsRulesService) Delete(ctx context.Context, ruleName string) (*AnalyticsRuleDeleteResponse, error) {
	u := fmt.Sprintf("/analytics/rules/%s", ruleName)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	res := &AnalyticsRuleDeleteResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
