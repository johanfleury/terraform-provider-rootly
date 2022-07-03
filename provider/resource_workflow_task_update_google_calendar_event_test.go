package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskUpdateGoogleCalendarEvent(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskUpdateGoogleCalendarEvent,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskUpdateGoogleCalendarEventUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskUpdateGoogleCalendarEvent = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_google_calendar_event" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		event_id = "test"
	}
}
`

const testAccResourceWorkflowTaskUpdateGoogleCalendarEventUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_google_calendar_event" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		event_id = "test"
	}
}
`
