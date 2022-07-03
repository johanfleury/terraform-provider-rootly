package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskSnapshotLookerLook(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskSnapshotLookerLook,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskSnapshotLookerLookUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskSnapshotLookerLook = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_snapshot_looker_look" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		dashboards {
						id = "foo"
						name = "bar"
					}
	}
}
`

const testAccResourceWorkflowTaskSnapshotLookerLookUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_snapshot_looker_look" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		dashboards {
						id = "foo"
						name = "bar"
					}
	}
}
`
