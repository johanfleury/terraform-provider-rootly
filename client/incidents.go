package client

import (
	"reflect"
	
	"github.com/pkg/errors"
	"github.com/google/jsonapi"
	rootlygo "github.com/rootlyhq/terraform-provider-rootly/schema"
)

type Incident struct {
	ID string `jsonapi:"primary,incidents"`
	Title string `jsonapi:"attr,title,omitempty"`
  Kind string `jsonapi:"attr,kind,omitempty"`
  Slug string `jsonapi:"attr,slug,omitempty"`
  ParentIncidentId string `jsonapi:"attr,parent_incident_id,omitempty"`
  Summary string `jsonapi:"attr,summary,omitempty"`
  Private *bool `jsonapi:"attr,private,omitempty"`
  Severity map[string]interface{} `jsonapi:"attr,severity,omitempty"`
  Environments []interface{} `jsonapi:"attr,environments,omitempty"`
  IncidentTypes []interface{} `jsonapi:"attr,incident_types,omitempty"`
  Services []interface{} `jsonapi:"attr,services,omitempty"`
  Functionalities []interface{} `jsonapi:"attr,functionalities,omitempty"`
  Groups []interface{} `jsonapi:"attr,groups,omitempty"`
  Labels map[string]interface{} `jsonapi:"attr,labels,omitempty"`
  SlackChannelId string `jsonapi:"attr,slack_channel_id,omitempty"`
  SlackChannelName string `jsonapi:"attr,slack_channel_name,omitempty"`
  SlackChannelUrl string `jsonapi:"attr,slack_channel_url,omitempty"`
  MitigationMessage string `jsonapi:"attr,mitigation_message,omitempty"`
  ResolutionMessage string `jsonapi:"attr,resolution_message,omitempty"`
  CancellationMessage string `jsonapi:"attr,cancellation_message,omitempty"`
  ScheduledFor string `jsonapi:"attr,scheduled_for,omitempty"`
  ScheduledUntil string `jsonapi:"attr,scheduled_until,omitempty"`
  InTriageAt string `jsonapi:"attr,in_triage_at,omitempty"`
  StartedAt string `jsonapi:"attr,started_at,omitempty"`
  DetectedAt string `jsonapi:"attr,detected_at,omitempty"`
  AcknowledgedAt string `jsonapi:"attr,acknowledged_at,omitempty"`
  MitigatedAt string `jsonapi:"attr,mitigated_at,omitempty"`
  ResolvedAt string `jsonapi:"attr,resolved_at,omitempty"`
  CancelledAt string `jsonapi:"attr,cancelled_at,omitempty"`
}

func (c *Client) ListIncidents(params *rootlygo.ListIncidentsParams) ([]interface{}, error) {
	req, err := rootlygo.NewListIncidentsRequest(c.Rootly.Server, params)
	if err != nil {
		return nil, errors.Errorf("Error building request: %s", err.Error())
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Errorf("Failed to make request: %s", err.Error())
	}

	incidents, err := jsonapi.UnmarshalManyPayload(resp.Body, reflect.TypeOf(new(Incident)))
	if err != nil {
		return nil, errors.Errorf("Error unmarshaling: %s", err.Error())
	}

	return incidents, nil
}

