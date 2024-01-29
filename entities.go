package gosmartis

import (
	"encoding/json"
	"strings"
	"time"
)

type AttributionModel int

const (
	AttributionModelLastClick              AttributionModel = 1  // Последнее касание
	AttributionModelFirstClick             AttributionModel = 2  // Первое касание
	AttributionModelLinear                 AttributionModel = 3  // Линейное распределение
	AttributionModelByPosition             AttributionModel = 4  // На основе позиции
	AttributionModelFirstCommunication     AttributionModel = 5  // Первое обращение
	AttributionModelLinearByCommunication  AttributionModel = 6  // Линейное распределение на обращениях
	AttributionModelLinearWithPostview     AttributionModel = 10 // Линейное распределение с учетом post-view
	AttributionModelLastClickWithPostview  AttributionModel = 15 // Последнее касание с учетом post-view
	AttributionModelFirstClickWithPostview AttributionModel = 16 // Первое касание с учетом post-view
	AttributionModelNotFirstNotLastClick   AttributionModel = 17 // Не первое и не последнее
	AttributionModelLastCommunication      AttributionModel = 22 // Последнее обращение
	AttributionModelByPositionWithPostview AttributionModel = 23 // На основе позиции с учетом post-view
)

type GroupBy string

const (
	GroupByAd        GroupBy = "ad_id"
	GroupByDay       GroupBy = "day"
	GroupByPlacement GroupBy = "placement_id"
	GroupByCampaign  GroupBy = "campaigns"
	GroupByObject    GroupBy = "smartis_object"
)

type FilterCategory int

const (
	FilterBySmartisID FilterCategory = 7071
	FilterByChannel   FilterCategory = 1222
	FilterByPlacement FilterCategory = 1223
)

type TypeReport string

const (
	TypeReportRaw        TypeReport = "raw"
	TypeReportAggregated TypeReport = "aggregated"
)

type GetChannelsResponse struct {
	Items []Channel `json:"channels"`
}

type Channel struct {
	ID                  int         `json:"id"`
	Title               string      `json:"title"`
	Name                *string     `json:"name"`
	IsActive            bool        `json:"isActive"`
	IsVisible           bool        `json:"isVisible"`
	IsDefaultForChannel bool        `json:"is_default_for_channel"`
	ParentChannelID     int         `json:"parent_channel_id"`
	NumLevel            int         `json:"num_level"`
	CatID               string      `json:"cat_id"`
	ServiceID           interface{} `json:"service_id"`
	ClientID            int         `json:"client_id"`
	GroupingID          int         `json:"grouping_id"`
	ClassData           interface{} `json:"classData"`
	GetDataMethod       interface{} `json:"getDataMethod"`
	DateCreate          string      `json:"date_create"`
	Sort                int         `json:"sort"`
	CreatedAt           *string     `json:"created_at"`
	UpdatedAt           *string     `json:"updated_at"`
	DeletedAt           interface{} `json:"deleted_at"`
	CategoryTitle       *string     `json:"category_title"`
}

type GetPlacementsResponse struct {
	Items []Placement `json:"placements"`
}

type Placement struct {
	ID                  int         `json:"id"`
	Title               string      `json:"title"`
	Name                *string     `json:"name"`
	IsActive            bool        `json:"isActive"`
	IsVisible           bool        `json:"isVisible"`
	IsDefaultForChannel bool        `json:"is_default_for_channel"`
	ParentChannelID     int         `json:"parent_channel_id"`
	NumLevel            int         `json:"num_level"`
	CatID               string      `json:"cat_id"`
	ServiceID           *int        `json:"service_id"`
	ClientID            int         `json:"client_id"`
	GroupingID          int         `json:"grouping_id"`
	ClassData           *string     `json:"classData"`
	GetDataMethod       interface{} `json:"getDataMethod"`
	DateCreate          string      `json:"date_create"`
	Sort                int         `json:"sort"`
	CreatedAt           *string     `json:"created_at"`
	UpdatedAt           *string     `json:"updated_at"`
	DeletedAt           interface{} `json:"deleted_at"`
	ChannelID           int         `json:"channel_id"`
	Channel             struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		ChannelID int    `json:"channel_id"`
	} `json:"channel"`
}

type GetCampaignsResponse struct {
	Items []Campaign `json:"campaigns"`
}

type Campaign struct {
	Id          int     `json:"id"`
	PlacementId int     `json:"placement_id"`
	Title       string  `json:"title"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type GetKeywordsResponse struct {
	Items []Keyword `json:"keywords"`
}

type Keyword struct {
	ID      int    `json:"id"`
	Keyword string `json:"keyword"`
}

type GetCrmCustomFieldsResponse struct {
	Items []CrmCustomField `json:"crmCustomFields"`
}

type CrmCustomField struct {
	ID                int    `json:"id"`
	CRMAccountId      int    `json:"crm_account_id"`
	ElementTypeId     int    `json:"element_type_id"`
	CustomFieldTitle  string `json:"custom_field_title"`
	FieldTypeId       int    `json:"field_type_id"`
	IsMultiple        int    `json:"is_multiple"`
	GroupId           int    `json:"group_id"`
	Description       string `json:"description"`
	Status            int    `json:"status"`
	IsFilter          int    `json:"is_filter"`
	FilterParamId     int    `json:"filter_param_id"`
	DefaultVisibility int    `json:"default_visibility"`
}

type GetCrmCustomFieldGroupsResponse struct {
	Items []CrmCustomFieldGroup `json:"crmCustomFieldGroups"`
}

type CrmCustomFieldGroup struct {
	ID                int    `json:"id"`
	Title             string `json:"title"`
	CRMAccountId      int    `json:"crm_account_id"`
	DefaultVisibility int    `json:"default_visibility"`
	Sort              int    `json:"sort"`
}

type GetAdsResponse struct {
	Items []Ad `json:"ads"`
}

type Ad struct {
	Id                 int         `json:"id"`
	ExternalId         string      `json:"external_id"`
	PlacementId        int         `json:"placement_id"`
	CampaignId         int         `json:"campaign_id"`
	ExternalCampaignId string      `json:"external_campaign_id"`
	Type               string      `json:"type"`
	Title              string      `json:"title"`
	Text               string      `json:"text"`
	Text1              string      `json:"text1"`
	Text2              *string     `json:"text2"`
	PreviewUrl         interface{} `json:"preview_url"`
	Href               interface{} `json:"href"`
	Device             interface{} `json:"device"`
	CreatedAt          string      `json:"created_at"`
}

type GetProjectsResponse struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Id                   int    `json:"id"`
	Project              string `json:"project"`
	Title                string `json:"title"`
	CreatedAt            int    `json:"created_at"`
	IsActive             int    `json:"is_active"`
	IsSuperObject        int    `json:"is_super_object"`
	CanGroupingByObjects int    `json:"can_grouping_by_objects"`
	ProjectFields        []struct {
		Value string `json:"value"`
		Title string `json:"title"`
	} `json:"project_fields"`
}

type GetMetricsResponse struct {
	Metrics []Metric `json:"metrics"`
}

type Metric struct {
	Id                 int     `json:"id"`
	Code               string  `json:"code"`
	Title              string  `json:"title"`
	Description        *string `json:"description"`
	CategoryId         int     `json:"category_id"`
	CategoryTitle      string  `json:"category_title"`
	CategorySort       int     `json:"category_sort"`
	IsSystem           int     `json:"is_system"`
	MParent            int     `json:"m_parent"`
	ServiceId          int     `json:"service_id"`
	IsGroup            *int    `json:"is_group"`
	Formule            *string `json:"formule"`
	Calculate          string  `json:"calculate"`
	DateCreate         int     `json:"date_create"`
	EnableOriginalData int     `json:"enable_original_data"`
}

type GetGroupingsResponse struct {
	Groupings []Grouping
}

type Grouping struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Code     string `json:"code"`
	IsSystem bool   `json:"is_system"`
	Sort     int    `json:"sort"`
	ClientId int    `json:"client_id"`
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
		var isSystem = v.IsSystem == 1
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

type AttributionSmartis struct {
	About    string `json:"about"`
	ID       int    `json:"id"`
	IsSystem bool   `json:"is_system"`
	Title    string `json:"title"`
}

type GetReportsResponse struct {
	Reports  map[string]interface{} `json:"reports"`
	MetaInfo struct {
		WorkTime int `json:"worktime"`
	} `json:"metaInfo"`
	Warnings []interface{} `json:"warnings"`
}

type Attribution struct {
	ModelID    AttributionModel `json:"model_id"`
	Period     int              `json:"period"`
	WithDirect bool             `json:"with_direct"`
}

type Filter struct {
	Name     string `json:"name"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Payload struct {
	Project      string
	Metrics      []string
	DateTimeFrom time.Time
	DateTimeTo   time.Time
	GroupBy      GroupBy
	TypeReport   TypeReport
	Filters      []Filter
	Fields       []string
	Attribution  Attribution
}

func (p *Payload) convert() map[string]interface{} {
	payload := map[string]interface{}{
		"project":      p.Project,
		"metrics":      strings.Join(p.Metrics, ";"),
		"datetimeFrom": p.DateTimeFrom.Format("2006-01-02"),
		"datetimeTo":   p.DateTimeTo.Format("2006-01-02"),
		"groupBy":      p.GroupBy,
		"type":         p.TypeReport,
		"attribution":  p.Attribution,
	}

	if len(p.Filters) != 0 {
		payload["filters"] = p.Filters
	}

	if len(p.Fields) != 0 {
		payload["fields"] = p.Fields
	}

	return payload
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
