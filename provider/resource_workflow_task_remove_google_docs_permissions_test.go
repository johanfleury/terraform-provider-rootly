package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskRemoveGoogleDocsPermissions(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskRemoveGoogleDocsPermissions,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskRemoveGoogleDocsPermissionsUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskRemoveGoogleDocsPermissions = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_remove_google_docs_permissions" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		file_id = "test"
attribute_to_query_by = "type"
value = "test"
	}
}
`

const testAccResourceWorkflowTaskRemoveGoogleDocsPermissionsUpdate = `
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_remove_google_docs_permissions" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		file_id = "test"
attribute_to_query_by = "type"
value = "test"
	}
}
`
