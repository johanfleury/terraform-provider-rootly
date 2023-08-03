package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskInviteToSlackChannel(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep {
			{
				Config: testAccResourceWorkflowTaskInviteToSlackChannel,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskInviteToSlackChannelUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskInviteToSlackChannel = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_invite_to_slack_channel" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		channel = {
					id = "foo"
					name = "bar"
				}
	}
}
`

const testAccResourceWorkflowTaskInviteToSlackChannelUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_invite_to_slack_channel" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		channel = {
					id = "foo"
					name = "bar"
				}
	}
}
`
