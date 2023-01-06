package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskTweetTwitterMessage(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskTweetTwitterMessage,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskTweetTwitterMessageUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskTweetTwitterMessage = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_tweet_twitter_message" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		message = "test"
	}
}
`

const testAccResourceWorkflowTaskTweetTwitterMessageUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_tweet_twitter_message" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		message = "test"
	}
}
`
