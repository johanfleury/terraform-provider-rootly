package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskPagePagerdutyOnCallResponders(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskPagePagerdutyOnCallResponders,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskPagePagerdutyOnCallRespondersUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskPagePagerdutyOnCallResponders = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_page_pagerduty_on_call_responders" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		service = {
					id = "foo"
					name = "bar"
				}
	}
}
`

const testAccResourceWorkflowTaskPagePagerdutyOnCallRespondersUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_page_pagerduty_on_call_responders" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		service = {
					id = "foo"
					name = "bar"
				}
	}
}
`
