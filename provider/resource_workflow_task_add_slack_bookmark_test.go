package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskAddSlackBookmark(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep {
			{
				Config: testAccResourceWorkflowTaskAddSlackBookmark,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskAddSlackBookmarkUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskAddSlackBookmark = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_add_slack_bookmark" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		channel = {
					id = "foo"
					name = "bar"
				}
	}
}
`

const testAccResourceWorkflowTaskAddSlackBookmarkUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_add_slack_bookmark" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		channel = {
					id = "foo"
					name = "bar"
				}
	}
}
`
