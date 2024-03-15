package gosmartis

import "encoding/json"

type GetChannelsResponse struct {
	Items []Channel `json:"channels"`
}

type GetPlacementsResponse struct {
	Items []Placement `json:"placements"`
}

type GetCampaignsResponse struct {
	Items []Campaign `json:"campaigns"`
}

type GetKeywordsResponse struct {
	Items []Keyword `json:"keywords"`
}

type GetCrmCustomFieldsResponse struct {
	Items []CrmCustomField `json:"crmCustomFields"`
}

type GetCrmCustomFieldGroupsResponse struct {
	Items []CrmCustomFieldGroup `json:"crmCustomFieldGroups"`
}

type GetAdsResponse struct {
	Items []Ad `json:"ads"`
}

type GetProjectsResponse struct {
	Projects []Project `json:"projects"`
}

type GetMetricsResponse struct {
	Metrics []Metric `json:"metrics"`
}

type GetGroupingsResponse struct {
	Groupings []Grouping
}

func (g *GetGroupingsResponse) UnmarshalJSON(b []byte) error {
	var data struct {
		Groupings map[string]struct {
			ID       int    `json:"id"`
			Title    string `json:"title"`
			Code     string `json:"code"`
			IsSystem bool   `json:"is_system"`
			Sort     int    `json:"sort"`
			ClientID int    `json:"client_id"`
		} `json:"groupings"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	groupings := make([]Grouping, 0, len(data.Groupings))
	for _, v := range data.Groupings {
		groupings = append(groupings, Grouping{
			Id:       v.ID,
			Title:    v.Title,
			Code:     v.Code,
			IsSystem: v.IsSystem,
			Sort:     v.Sort,
			ClientId: v.ClientID,
		})
	}

	g.Groupings = groupings

	return nil
}

type GetAttributionsResponse struct {
	Attributions []AttributionSmartis
}

func (g *GetAttributionsResponse) UnmarshalJSON(b []byte) error {
	var data struct {
		ModelAttributions map[string]struct {
			ID       int    `json:"id"`
			IsSystem int    `json:"is_system"`
			About    string `json:"about"`
			Title    string `json:"title"`
		} `json:"modelAttributions"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	attributions := make([]AttributionSmartis, 0, len(data.ModelAttributions))

	for _, v := range data.ModelAttributions {
		isSystem := v.IsSystem == 1
		attributions = append(attributions, AttributionSmartis{
			About:    v.About,
			ID:       v.ID,
			IsSystem: isSystem,
			Title:    v.Title,
		})
	}

	g.Attributions = attributions

	return nil
}

type GetReportsResponse struct {
	Reports  map[string]interface{} `json:"reports"`
	MetaInfo struct {
		WorkTime int `json:"worktime"`
	} `json:"metaInfo"`
	Warnings []interface{} `json:"warnings"`
}
