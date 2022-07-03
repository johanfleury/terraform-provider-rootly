package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskCreateLinearSubtaskIssue(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskCreateLinearSubtaskIssue,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskCreateLinearSubtaskIssueUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskCreateLinearSubtaskIssue = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_linear_subtask_issue" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		parent_issue_id = "test"
title = "test"
state = {
					id = "foo"
					name = "bar"
				}
	}
}
`

const testAccResourceWorkflowTaskCreateLinearSubtaskIssueUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_linear_subtask_issue" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		parent_issue_id = "test"
title = "test"
state = {
					id = "foo"
					name = "bar"
				}
	}
}
`
