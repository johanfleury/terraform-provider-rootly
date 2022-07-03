package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskCreateTrelloCard(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskCreateTrelloCard,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskCreateTrelloCardUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskCreateTrelloCard = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_trello_card" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		title = "test"
board = {
					id = "foo"
					name = "bar"
				}
list = {
					id = "foo"
					name = "bar"
				}
	}
}
`

const testAccResourceWorkflowTaskCreateTrelloCardUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_create_trello_card" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		title = "test"
board = {
					id = "foo"
					name = "bar"
				}
list = {
					id = "foo"
					name = "bar"
				}
	}
}
`
