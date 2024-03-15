package gosmartis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	baseURL = "https://my.smartis.bi/api/"
)

type Client struct {
	APIKey     string
	CRMToken   string
	HOST       string
	HTTPClient *http.Client
}

func NewClient(apiKey, crmToken string, httpClient *http.Client) *Client {
	return &Client{
		APIKey:     apiKey,
		CRMToken:   crmToken,
		HOST:       baseURL,
		HTTPClient: httpClient,
	}
}

type Method struct {
	endpoint string
	plural   string
}

func (m *Method) getURL() string {
	return fmt.Sprintf("%s%s", baseURL, m.endpoint)
}

var (
	GetReports              = Method{plural: "reports", endpoint: "reports/getReport"}
	GetProjects             = Method{plural: "projects", endpoint: "projects/get"}
	GetMetrics              = Method{plural: "metrics", endpoint: "metrics/get"}
	GetGroupings            = Method{plural: "groupings", endpoint: "reports/getGroupings"}
	GetAttributions         = Method{plural: "attributions", endpoint: "reports/getModelAttributions"}
	GetChannels             = Method{plural: "", endpoint: "reports/getChannels"}
	GetPlacements           = Method{plural: "", endpoint: "reports/getPlacements"}
	GetCampaigns            = Method{plural: "", endpoint: "reports/getCampaigns"}
	GetAds                  = Method{plural: "", endpoint: "reports/getAds"}
	GetKeywords             = Method{plural: "", endpoint: "reports/getKeywords"}
	GetCRMCustomFields      = Method{plural: "crmCustomFields", endpoint: "crm/crmCustomField/get"}
	GetCRMCustomFieldGroups = Method{plural: "crmCustomFieldGroups", endpoint: "crm/crmCustomFieldGroup/get"}
)

func statusCodeHandler(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	switch resp.StatusCode {
	case http.StatusInternalServerError:
		return errInternalError
	case http.StatusUnauthorized:
		return errUnauthorized
	default:
		var errMsg APIError

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(respBody, &errMsg)
		if err != nil {
			return fmt.Errorf("unknown error: %d", resp.StatusCode)
		}

		return &errMsg
	}
}

func (c *Client) doRequest(ctx context.Context, method Method, data interface{}) (*http.Response, error) {
	header := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", c.APIKey),
		"Content-Type":  "application/json",
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	if data != nil {
		err := enc.Encode(data)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(http.MethodPost, method.getURL(), &buf)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = statusCodeHandler(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetReport(ctx context.Context, payload Payload) ([]*Report, error) {
	pay := payload.convert()

	resp, err := c.doRequest(ctx, GetReports, pay)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	data := GetReportsResponse{}

	err = json.Unmarshal(respBody, &data)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	if len(data.Reports) == 0 {
		return nil, fmt.Errorf("no reports data")
	}

	reports := make([]*Report, 0, len(data.Reports))

	for k, v := range data.Reports {

		report, err := newReport(k, v)
		if err != nil {
			return nil, err
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func newReport(metric string, data interface{}) (*Report, error) {
	prepare, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid report data")
	}
	report := Report{
		Metric:      metric,
		RowsMassive: make([]Row, 0, len(prepare)),
	}

	for _, row := range prepare {
		rowData := make([]*Cell, 0, len(row.(map[string]interface{})))

		prepareRow, ok := row.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid report data")
		}

		for key, value := range prepareRow {
			cell := Cell{
				ColumnID: key,
				Value:    value,
				Name:     "",
			}
			cell.initType()
			rowData = append(rowData, &cell)
		}

		report.RowsMassive = append(report.RowsMassive, rowData)
	}

	return &report, nil
}

// GetProjects retrieves a list of projects using the provided context.
// It returns a slice of Project and an error.
func (c *Client) GetProjects(ctx context.Context) ([]Project, error) {
	resp, err := c.doRequest(ctx, GetProjects, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result GetProjectsResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Projects, nil
}

// GetMetrics retrieves a list of metrics using the provided context.
// It returns a slice of Metric and an error.
func (c *Client) GetMetrics(ctx context.Context) ([]Metric, error) {
	resp, err := c.doRequest(ctx, GetMetrics, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result GetMetricsResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Metrics, nil
}

// GetGroupings retrieves a list of groupings using the provided context.
// It returns a slice of Grouping and an error.
func (c *Client) GetGroupings(ctx context.Context) ([]Grouping, error) {
	resp, err := c.doRequest(ctx, GetGroupings, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result GetGroupingsResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Groupings, nil
}

// GetAttributions retrieves a list of attributions using the provided context.
// It returns a slice of AttributionSmartis and an error.
func (c *Client) GetAttributions(ctx context.Context) ([]AttributionSmartis, error) {
	resp, err := c.doRequest(ctx, GetAttributions, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result GetAttributionsResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Attributions, nil
}

// GetChannels retrieves a list of channels using the provided context.
// It returns a slice of Channel and an error.
func (c *Client) GetChannels(ctx context.Context) ([]Channel, error) {
	resp, err := c.doRequest(ctx, GetChannels, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result GetChannelsResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}

// GetPlacements retrieves a list of placements using the provided context.
// It returns a slice of Placement and an error.
func (c *Client) GetPlacements(ctx context.Context) ([]Placement, error) {
	resp, err := c.doRequest(ctx, GetPlacements, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result GetPlacementsResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Items, err
}

// GetCampaigns retrieves a list of campaigns using the provided context.
// It returns a slice of Campaign and an error.
func (c *Client) GetCampaigns(ctx context.Context, ids []int) ([]Campaign, error) {
	data := map[string]interface{}{
		"ids": ids,
	}

	resp, err := c.doRequest(ctx, GetCampaigns, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result GetCampaignsResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}

// GetAds retrieves a list of ads using the provided context.
// It returns a slice of Ad and an error.
func (c *Client) GetAds(ctx context.Context, ids []int) ([]Ad, error) {
	data := map[string]interface{}{
		"ids": ids,
	}

	resp, err := c.doRequest(ctx, GetAds, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result GetAdsResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}

// GetKeywords retrieves a list of keywords using the provided context.
// It returns a slice of Keyword and an error.
func (c *Client) GetKeywords(ctx context.Context, ids []int) ([]Keyword, error) {
	data := map[string]interface{}{
		"ids": ids,
	}

	resp, err := c.doRequest(ctx, GetKeywords, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result GetKeywordsResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Items, err
}

// GetCRMCustomFields retrieves a list of custom fields using the provided context.
// It returns a slice of CrmCustomField and an error.
func (c *Client) GetCRMCustomFields(ctx context.Context, ids []int) ([]CrmCustomField, error) {
	data := map[string]interface{}{
		"ids":               ids,
		"smartis_crm_token": c.CRMToken,
	}

	resp, err := c.doRequest(ctx, GetCRMCustomFields, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result GetCrmCustomFieldsResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Items, err
}

// GetCRMCustomFieldGroups retrieves a list of custom field groups using the provided context.
// It returns a slice of CrmCustomFieldGroup and an error.
func (c *Client) GetCRMCustomFieldGroups(ctx context.Context, ids []int) ([]CrmCustomFieldGroup, error) {
	data := map[string]interface{}{
		"ids":               ids,
		"smartis_crm_token": c.CRMToken,
	}

	resp, err := c.doRequest(ctx, GetCRMCustomFieldGroups, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result GetCrmCustomFieldGroupsResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Items, err
}
