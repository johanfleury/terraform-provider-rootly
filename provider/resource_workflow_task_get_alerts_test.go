package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskGetAlerts(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep {
			{
				Config: testAccResourceWorkflowTaskGetAlerts,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskGetAlertsUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskGetAlerts = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_get_alerts" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		past_duration = "1 hour"
	}
}
`

const testAccResourceWorkflowTaskGetAlertsUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_get_alerts" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		past_duration = "1 hour"
	}
}
`
