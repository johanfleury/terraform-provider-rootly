package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskAddRole(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep {
			{
				Config: testAccResourceWorkflowTaskAddRole,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskAddRoleUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskAddRole = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_add_role" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		incident_role_id = "test"
	}
}
`

const testAccResourceWorkflowTaskAddRoleUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_add_role" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		incident_role_id = "test"
	}
}
`
