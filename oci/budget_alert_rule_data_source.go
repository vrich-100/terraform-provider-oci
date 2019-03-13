// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_budget "github.com/oracle/oci-go-sdk/budget"
)

func BudgetAlertRuleDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularBudgetAlertRule,
		Schema: map[string]*schema.Schema{
			"alert_rule_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"budget_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recipients": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"threshold": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"threshold_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularBudgetAlertRule(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetAlertRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).budgetClient

	return ReadResource(sync)
}

type BudgetAlertRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_budget.BudgetClient
	Res    *oci_budget.GetAlertRuleResponse
}

func (s *BudgetAlertRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BudgetAlertRuleDataSourceCrud) Get() error {
	request := oci_budget.GetAlertRuleRequest{}

	if alertRuleId, ok := s.D.GetOkExists("alert_rule_id"); ok {
		tmp := alertRuleId.(string)
		request.AlertRuleId = &tmp
	}

	if budgetId, ok := s.D.GetOkExists("budget_id"); ok {
		tmp := budgetId.(string)
		request.BudgetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "budget")

	response, err := s.Client.GetAlertRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BudgetAlertRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Message != nil {
		s.D.Set("message", *s.Res.Message)
	}

	if s.Res.Recipients != nil {
		s.D.Set("recipients", *s.Res.Recipients)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Threshold != nil {
		s.D.Set("threshold", *s.Res.Threshold)
	}

	s.D.Set("threshold_type", s.Res.ThresholdType)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}