package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskUpdateTrelloCard(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskUpdateTrelloCard,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskUpdateTrelloCardUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskUpdateTrelloCard = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_trello_card" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		card_id = "test"
archivation = {
					id = "foo"
					name = "bar"
				}
	}
}
`

const testAccResourceWorkflowTaskUpdateTrelloCardUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_trello_card" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		card_id = "test"
archivation = {
					id = "foo"
					name = "bar"
				}
	}
}
`
