package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskUpdateJiraIssue(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskUpdateJiraIssue,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskUpdateJiraIssueUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskUpdateJiraIssue = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_jira_issue" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		issue_id = "test"
project_key = "test"
	}
}
`

const testAccResourceWorkflowTaskUpdateJiraIssueUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_jira_issue" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		issue_id = "test"
project_key = "test"
	}
}
`
