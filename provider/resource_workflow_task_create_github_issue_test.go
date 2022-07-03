package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskCreateGithubIssue(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskCreateGithubIssue,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskCreateGithubIssueUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskCreateGithubIssue = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_github_issue" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		repository = {
					id = "foo"
					name = "bar"
				}
title = "test"
	}
}
`

const testAccResourceWorkflowTaskCreateGithubIssueUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_github_issue" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		repository = {
					id = "foo"
					name = "bar"
				}
title = "test"
	}
}
`
