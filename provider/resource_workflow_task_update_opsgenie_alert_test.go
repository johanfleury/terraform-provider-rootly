package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskUpdateOpsgenieAlert(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskUpdateOpsgenieAlert,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskUpdateOpsgenieAlertUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskUpdateOpsgenieAlert = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_opsgenie_alert" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		alert_id = "test"
priority = "P1"
completion = {
					id = "foo"
					name = "bar"
				}
	}
}
`

const testAccResourceWorkflowTaskUpdateOpsgenieAlertUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_opsgenie_alert" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		alert_id = "test"
priority = "P1"
completion = {
					id = "foo"
					name = "bar"
				}
	}
}
`
