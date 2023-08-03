package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskUpdateNotionPage(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep {
			{
				Config: testAccResourceWorkflowTaskUpdateNotionPage,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskUpdateNotionPageUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskUpdateNotionPage = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_notion_page" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		file_id = "test"
	}
}
`

const testAccResourceWorkflowTaskUpdateNotionPageUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_notion_page" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		file_id = "test"
	}
}
`
