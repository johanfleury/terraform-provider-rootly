package client

import (
	"reflect"
	
	"github.com/pkg/errors"
	"github.com/google/jsonapi"
	rootlygo "github.com/rootlyhq/terraform-provider-rootly/schema"
)

type IncidentPostMortem struct {
	ID string `jsonapi:"primary,incident_post_mortems"`
	Title string `jsonapi:"attr,title,omitempty"`
  Content string `jsonapi:"attr,content,omitempty"`
  Status string `jsonapi:"attr,status,omitempty"`
  StartedAt string `jsonapi:"attr,started_at,omitempty"`
  MitigatedAt string `jsonapi:"attr,mitigated_at,omitempty"`
  ResolvedAt string `jsonapi:"attr,resolved_at,omitempty"`
  ShowTimeline *bool `jsonapi:"attr,show_timeline,omitempty"`
  ShowTimelineTrail *bool `jsonapi:"attr,show_timeline_trail,omitempty"`
  ShowTimelineGenius *bool `jsonapi:"attr,show_timeline_genius,omitempty"`
  ShowTimelineTasks *bool `jsonapi:"attr,show_timeline_tasks,omitempty"`
  ShowTimelineActionItems *bool `jsonapi:"attr,show_timeline_action_items,omitempty"`
  ShowTimelineOrder string `jsonapi:"attr,show_timeline_order,omitempty"`
  ShowServicesImpacted *bool `jsonapi:"attr,show_services_impacted,omitempty"`
  ShowFunctionalitiesImpacted *bool `jsonapi:"attr,show_functionalities_impacted,omitempty"`
  ShowGroupsImpacted *bool `jsonapi:"attr,show_groups_impacted,omitempty"`
  ShowAlertsAttached *bool `jsonapi:"attr,show_alerts_attached,omitempty"`
  Url string `jsonapi:"attr,url,omitempty"`
}

func (c *Client) ListIncidentPostMortems(params *rootlygo.ListIncidentPostMortemsParams) ([]interface{}, error) {
	req, err := rootlygo.NewListIncidentPostMortemsRequest(c.Rootly.Server, params)
	if err != nil {
		return nil, errors.Errorf("Error building request: %s", err.Error())
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Errorf("Failed to make request: %s", err.Error())
	}

	incident_post_mortems, err := jsonapi.UnmarshalManyPayload(resp.Body, reflect.TypeOf(new(IncidentPostMortem)))
	if err != nil {
		return nil, errors.Errorf("Error unmarshaling: %s", err.Error())
	}

	return incident_post_mortems, nil
}

