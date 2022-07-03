package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskSnapshotNewRelicGraph(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskSnapshotNewRelicGraph,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskSnapshotNewRelicGraphUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskSnapshotNewRelicGraph = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_snapshot_new_relic_graph" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		metric_query = "test"
metric_type = "APDEX"
	}
}
`

const testAccResourceWorkflowTaskSnapshotNewRelicGraphUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_snapshot_new_relic_graph" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		metric_query = "test"
metric_type = "APDEX"
	}
}
`
