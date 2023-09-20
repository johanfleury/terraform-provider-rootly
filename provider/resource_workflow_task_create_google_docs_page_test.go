package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskCreateGoogleDocsPage(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep {
			{
				Config: testAccResourceWorkflowTaskCreateGoogleDocsPage,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskCreateGoogleDocsPageUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskCreateGoogleDocsPage = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_google_docs_page" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		title = "test"
	}
}
`

const testAccResourceWorkflowTaskCreateGoogleDocsPageUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_google_docs_page" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		title = "test"
	}
}
`
