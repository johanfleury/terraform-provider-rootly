package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTaskUpdateAirtableTableRecord(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkflowTaskUpdateAirtableTableRecord,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTaskUpdateAirtableTableRecordUpdate,
			},
		},
	})
}

const testAccResourceWorkflowTaskUpdateAirtableTableRecord = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_airtable_table_record" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		base_key = "test"
table_name = "test"
record_id = "test"
	}
}
`

const testAccResourceWorkflowTaskUpdateAirtableTableRecordUpdate = `
resource "rootly_workflow_incident" "foo" {
  name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_update_airtable_table_record" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		base_key = "test"
table_name = "test"
record_id = "test"
	}
}
`
