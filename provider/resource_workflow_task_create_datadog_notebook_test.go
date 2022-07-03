package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskCreateDatadogNotebook(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskCreateDatadogNotebook,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskCreateDatadogNotebookUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskCreateDatadogNotebook = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_datadog_notebook" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		title = "test"
kind = "postmortem"
	}
}
`

const testAccResourceWorkflowTaskCreateDatadogNotebookUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_datadog_notebook" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		title = "test"
kind = "postmortem"
	}
}
`
