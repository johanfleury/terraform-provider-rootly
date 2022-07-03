package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskCreateAsanaTask(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskCreateAsanaTask,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskCreateAsanaTaskUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskCreateAsanaTask = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_asana_task" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		workspace = {
					id = "foo"
					name = "bar"
				}
projects {
						id = "foo"
						name = "bar"
					}
title = "test"
completion = {
					id = "foo"
					name = "bar"
				}
	}
}
`

const testAccResourceWorkflowTaskCreateAsanaTaskUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_asana_task" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		workspace = {
					id = "foo"
					name = "bar"
				}
projects {
						id = "foo"
						name = "bar"
					}
title = "test"
completion = {
					id = "foo"
					name = "bar"
				}
	}
}
`
