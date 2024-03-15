package gosmartis

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type CellType string

const (
	CellTypeSystemField  CellType = "system_field"
	CellTypeField        CellType = "field"
	CellTypeFieldCfGroup CellType = "field_cf_group"
)

type Report struct {
	Metric            string
	RowsMassive       []Row
	RowsMapped        []RowMapped
	isMapped          bool
	isGotColumnsNames bool
}

func (r *Report) IsGotColumnsNames() bool {
	return r.isGotColumnsNames
}

func (r *Report) IsMapped() bool {
	return r.isMapped
}

func (r *Report) GetColumnsNames(ctx context.Context, client *Client) error {
	if client.CRMToken == "" {
		return fmt.Errorf("crm token is empty")
	}

	crmCustomFieldIDs := make(map[string][]*Cell)
	crmCustomFieldGroupIDS := make(map[string][]*Cell)
	for _, row := range r.RowsMassive {
		for _, cell := range row {
			if cell.Type == CellTypeField {
				crmCustomFieldIDs[cell.CleanID] = append(crmCustomFieldIDs[cell.CleanID], cell)
			} else if cell.Type == CellTypeFieldCfGroup {
				crmCustomFieldGroupIDS[cell.CleanID] = append(crmCustomFieldGroupIDS[cell.CleanID], cell)
			}
		}
	}

	if len(crmCustomFieldIDs) > 0 {
		crmCustomFieldIDSList := make([]int, 0, len(crmCustomFieldIDs))
		for id := range crmCustomFieldIDs {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}

			crmCustomFieldIDSList = append(crmCustomFieldIDSList, idInt)
		}
		customFieldNames, err := client.GetCRMCustomFields(ctx, crmCustomFieldIDSList)
		if err != nil {
			return err
		}

		for _, field := range customFieldNames {
			cells := crmCustomFieldIDs[strconv.Itoa(field.ID)]
			for _, cell := range cells {
				cell.Name = field.CustomFieldTitle
			}
		}
	}
	if len(crmCustomFieldGroupIDS) > 0 {
		crmCustomFieldGroupIDSList := make([]int, 0, len(crmCustomFieldGroupIDS))

		for id := range crmCustomFieldGroupIDS {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}

			crmCustomFieldGroupIDSList = append(crmCustomFieldGroupIDSList, idInt)
		}
		customFieldGroupNames, err := client.GetCRMCustomFieldGroups(context.Background(), crmCustomFieldGroupIDSList)
		if err != nil {
			return err
		}

		for _, field := range customFieldGroupNames {
			cells := crmCustomFieldGroupIDS[strconv.Itoa(field.ID)]
			for _, cell := range cells {
				cell.Name = field.Title
			}
		}
	}

	r.isGotColumnsNames = true

	return nil
}

func (r *Report) MapColumns() error {
	for _, row := range r.RowsMassive {
		mapped := make(map[string]*Cell, len(row))
		for _, cell := range row {
			mapped[cell.ColumnID] = cell
		}

		r.RowsMapped = append(r.RowsMapped, mapped)
	}

	r.isMapped = true

	return nil
}

type Row []*Cell

type RowMapped map[string]*Cell

type Cell struct {
	ColumnID string
	CleanID  string
	Value    interface{}
	Name     string
	Type     CellType
}

func (c *Cell) initType() {
	if strings.Contains(c.ColumnID, "field_") && !strings.Contains(c.ColumnID, "field_cf_group_") {
		c.Type = CellTypeField
		rawColumnID := strings.Replace(c.ColumnID, "field_", "", 1)
		c.CleanID = strings.Split(rawColumnID, "_")[0]

		return
	} else if strings.Contains(c.ColumnID, "field_cf_group_") {
		c.Type = CellTypeFieldCfGroup
		c.CleanID = c.ColumnID
		rawColumnID := strings.Replace(c.ColumnID, "field_cf_group_", "", 1)
		c.CleanID = strings.Split(rawColumnID, "_")[0]

		return
	}

	c.Type = CellTypeSystemField
}
