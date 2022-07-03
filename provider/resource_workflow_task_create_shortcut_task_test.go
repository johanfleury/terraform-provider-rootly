package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskCreateShortcutTask(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskCreateShortcutTask,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskCreateShortcutTaskUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskCreateShortcutTask = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_shortcut_task" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		parent_story_id = "test"
description = "test"
completion = {
					id = "foo"
					name = "bar"
				}
	}
}
`

const testAccResourceWorkflowTaskCreateShortcutTaskUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_shortcut_task" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		parent_story_id = "test"
description = "test"
completion = {
					id = "foo"
					name = "bar"
				}
	}
}
`
