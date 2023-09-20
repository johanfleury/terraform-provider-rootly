package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskPageOpsgenieOnCallResponders(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep {
			{
				Config: testAccResourceWorkflowTaskPageOpsgenieOnCallResponders,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskPageOpsgenieOnCallRespondersUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskPageOpsgenieOnCallResponders = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_page_opsgenie_on_call_responders" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		
	}
}
`

const testAccResourceWorkflowTaskPageOpsgenieOnCallRespondersUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_page_opsgenie_on_call_responders" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		
	}
}
`
