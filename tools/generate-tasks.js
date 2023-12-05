const fs = require("fs");
const path = require("path");

const excluded = [];

console.log(`Excluding task from generation:`, excluded);

module.exports = (swagger) => {
  const schema = swagger.components.schemas;
  return Object.keys(schema)
    .filter((key) => key.match(/_task_params/))
    .map((key) => {
      const task_name = key.replace("_task_params", "");
      const task_schema = schema[key];
      const task_name_camel = task_name
        .split("_")
        .map((p) => `${p[0].toUpperCase()}${p.slice(1)}`)
        .join("");
      if (!excluded.filter((k) => key.match(k)).length) {
        fs.writeFileSync(
          `./provider/resource_workflow_task_${task_name}.go`,
          genResourceFile(task_name, task_schema)
        );
      }
      return task_name;
    });
};

function genResourceFile(task_name, task_schema) {
  const task_name_camel = task_name
    .split("_")
    .map((p) => `${p[0].toUpperCase()}${p.slice(1)}`)
    .join("");
  const has_json_field =
    Object.keys(task_schema.properties).filter((key) => {
      const field_schema = task_schema.properties[key];
      return (
        field_schema.type === "string" &&
        field_schema.description &&
        field_schema.description.match(/JSON/)
      );
    }).length > 0;
  return `package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"context"
	"fmt"
	${
    has_json_field
      ? `
	"reflect"
  "encoding/json"
	`
      : ``
  }
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/rootlyhq/terraform-provider-rootly/client"
	"github.com/rootlyhq/terraform-provider-rootly/tools"
)

func resourceWorkflowTask${task_name_camel}() *schema.Resource {
	return &schema.Resource{
		Description: "Manages workflow ${task_name} task.",

		CreateContext: resourceWorkflowTask${task_name_camel}Create,
		ReadContext:   resourceWorkflowTask${task_name_camel}Read,
		UpdateContext: resourceWorkflowTask${task_name_camel}Update,
		DeleteContext: resourceWorkflowTask${task_name_camel}Delete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema {
			"workflow_id": {
				Description:  "The ID of the parent workflow",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
			},
			"name": {
				Description:  "Name of the workflow task",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
			},
			"position": {
				Description:  "The position of the workflow task (1 being top of list)",
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
			},
			"skip_on_failure": {
				Description:  "Skip workflow task if any failures",
				Type:         schema.TypeBool,
				Optional:     true,
				Default:      false,
			},
			"enabled": {
				Description:  "Enable/disable this workflow task",
				Type:         schema.TypeBool,
				Optional:     true,
				Default:      true,
			},
			"task_params": {
				Description: "The parameters for this workflow task.",
				Type: schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema {
						"task_type": &schema.Schema {
							Type: schema.TypeString,
							Optional: true,
							Default: "${task_name}",
							ValidateFunc: validation.StringInSlice([]string {
								"${task_name}",
							}, false),
						},
${Object.keys(task_schema.properties)
  .filter((k) => k !== "task_type")
  .map((key) =>
    genTaskSchemaProperty(
      key,
      task_schema.properties[key],
      task_schema.required
    )
  )
  .join("\n")}
					},
				},
			},
		},
	}
}

func resourceWorkflowTask${task_name_camel}Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	workflowId := d.Get("workflow_id").(string)
	name := d.Get("name").(string)
	position := d.Get("position").(int)
	skipOnFailure := tools.Bool(d.Get("skip_on_failure").(bool))
	enabled := tools.Bool(d.Get("enabled").(bool))
	taskParams := d.Get("task_params").([]interface{})[0].(map[string]interface{})

	tflog.Trace(ctx, fmt.Sprintf("Creating workflow task: %s", workflowId))

	s := &client.WorkflowTask{
		WorkflowId: workflowId,
		Name: name,
		Position: position,
		SkipOnFailure: skipOnFailure,
		Enabled: enabled,
		TaskParams: taskParams,
	}

	res, err := c.CreateWorkflowTask(s)
	if err != nil {
		return diag.Errorf("Error creating workflow task: %s", err.Error())
	}

	d.SetId(res.ID)
	tflog.Trace(ctx, fmt.Sprintf("created an workflow task resource: %v (%s)", workflowId, d.Id()))

	return resourceWorkflowTask${task_name_camel}Read(ctx, d, meta)
}

func resourceWorkflowTask${task_name_camel}Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading workflow task: %s", d.Id()))

	res, err := c.GetWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTask${task_name_camel} (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error reading workflow task: %s", d.Id())
	}

	d.Set("workflow_id", res.WorkflowId)
	d.Set("name", res.Name)
	d.Set("position", res.Position)
	d.Set("skip_on_failure", res.SkipOnFailure)
	d.Set("enabled", res.Enabled)
	tps := make([]interface{}, 1, 1)
	tps[0] = res.TaskParams
	d.Set("task_params", tps)

	return nil
}

func resourceWorkflowTask${task_name_camel}Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating workflow task: %s", d.Id()))

	workflowId := d.Get("workflow_id").(string)
	name := d.Get("name").(string)
	position := d.Get("position").(int)
	skipOnFailure := tools.Bool(d.Get("skip_on_failure").(bool))
	enabled := tools.Bool(d.Get("enabled").(bool))
	taskParams := d.Get("task_params").([]interface{})[0].(map[string]interface{})

	s := &client.WorkflowTask{
		WorkflowId: workflowId,
		Name: name,
		Position: position,
		SkipOnFailure: skipOnFailure,
		Enabled: enabled,
		TaskParams: taskParams,
	}

	tflog.Debug(ctx, fmt.Sprintf("adding value: %#v", s))
	_, err := c.UpdateWorkflowTask(d.Id(), s)
	if err != nil {
		return diag.Errorf("Error updating workflow task: %s", err.Error())
	}

	return resourceWorkflowTask${task_name_camel}Read(ctx, d, meta)
}

func resourceWorkflowTask${task_name_camel}Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting workflow task: %s", d.Id()))

	err := c.DeleteWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTask${task_name_camel} (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting workflow task: %s", err.Error())
	}

	d.SetId("")

	return nil
}
`;
}

function annotatedDescription(schema) {
  const description = (schema.description || "")
    .replace(/"/g, '\\"')
    .replace(/ex. .+$/, "");
  if (schema.enum) {
    return `${
      !!description ? `${description}. ` : ""
    }Value must be one of ${schema.enum
      .map((val) => `\`${val}\``)
      .join(", ")}.`;
  }
  if (schema.type === "array" && schema.items && schema.items.enum) {
    return `${
      !!description ? `${description}. ` : ""
    }Value must be one of ${schema.items.enum
      .map((val) => `\`${val}\``)
      .join(", ")}.`;
  }
  if (
    schema.type === "object" &&
    schema.properties.id &&
    schema.properties.name
  ) {
    return `Map must contain two fields, \`id\` and \`name\`. ${description}`;
  }
  if (schema.type === "boolean") {
    return `${
      !!description ? `${description}. ` : ""
    }Value must be one of true or false`;
  }
  return description;
}

function genTaskSchemaProperty(property_name, property_schema, required_props) {
  const isRequired =
    required_props && required_props.indexOf(property_name) !== -1;
  const isJSON =
    property_schema.type === "string" &&
    property_schema.description &&
    property_schema.description.match(/JSON/);
  let a = `						"${property_name}": &schema.Schema {
							Description: "${annotatedDescription(property_schema)}",
							Type: ${genTaskSchemaPropertyType(property_schema.type)},
							${isRequired ? "Required" : "Optional"}: true,`;
  if (isJSON) {
    a = `${a}
							DiffSuppressFunc: func(k, old string, new string, d *schema.ResourceData) bool {
								var oldJSONAsInterface, newJSONAsInterface interface{}
							
								if err := json.Unmarshal([]byte(old), &oldJSONAsInterface); err != nil {
									return false
								}

								if err := json.Unmarshal([]byte(new), &newJSONAsInterface); err != nil {
									return false
								}

								return reflect.DeepEqual(oldJSONAsInterface, newJSONAsInterface)
							},`;
    if (!isRequired) {
      a = `${a}
							Default: "{}",`;
    }
  }
  if (property_schema.enum) {
    if (!isRequired) {
      if (property_schema?.default) {
        a = `${a}
							Default: "${property_schema?.default}",`;
      } else {
        a = `${a}
							Default: nil,`;
      }
    }
    a = `${a}
							ValidateFunc: validation.StringInSlice([]string{
								${property_schema.enum.map((k) => `"${k}",`).join("\n")}
							}, false),`;
  }
  if (property_schema.type === "number") {
    if (!isRequired) {
      a = `${a}
							Default: nil,`;
    }
  }
  if (property_schema.type === "array") {
    if (property_schema.items.type === "string") {
      a = `${a}
							Elem: &schema.Schema {
								Type: schema.TypeString,
							},`;
    } else if (property_schema.items.type === "object") {
      a = `${a}
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema {
									"id": &schema.Schema {
										Type: schema.TypeString,
										Required: true,
									},
									"name": &schema.Schema {
										Type: schema.TypeString,
										Required: true,
									},
								},
							},`;
    }
  }
  return `${a}
						},`;
}

function genTaskSchemaPropertyType(type_) {
  switch (type_) {
    case "string":
      return "schema.TypeString";
    case "number":
      return "schema.TypeInt";
    case "array":
      return "schema.TypeList";
    case "object":
      return "schema.TypeMap";
    case "boolean":
      return "schema.TypeBool";
    default:
      return "schema.TypeString";
  }
}

function genResourceTestFile(task_name, task_schema) {
  const task_name_camel = task_name
    .split("_")
    .map((p) => `${p[0].toUpperCase()}${p.slice(1)}`)
    .join("");
  return `package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWorkflowTask${task_name_camel}(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep {
			{
				Config: testAccResourceWorkflowTask${task_name_camel},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rootly_workflow_incident.foo", "name", "test-workflow"),
				),
			},
			{
				Config: testAccResourceWorkflowTask${task_name_camel}Update,
			},
		},
	})
}

const testAccResourceWorkflowTask${task_name_camel} = \`
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_${task_name}" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		${genTestParams(task_name, task_schema).join("\n")}
	}
}
\`

const testAccResourceWorkflowTask${task_name_camel}Update = \`
resource "rootly_workflow_incident" "foo" {
  	name = "test-workflow"
	trigger_params {
		triggers = ["incident_updated"]
	}
}

resource "rootly_workflow_task_${task_name}" "foo" {
	workflow_id = rootly_workflow_incident.foo.id
	task_params {
		${genTestParams(task_name, task_schema).join("\n")}
	}
}
\`
`;
}

function genTestParams(task_name, task_schema) {
  const required = task_schema.required || [];
  return required.map((key) => {
    let val;

    if (task_schema.properties[key].example) {
      val = task_schema.properties[key].example;
    } else if (task_schema.properties[key].enum) {
      val = task_schema.properties[key].enum[0];
    } else {
      val = "test";
    }

    switch (task_schema.properties[key].type) {
      case "boolean":
        return `${key} = false`;
      case "string":
        return `${key} = "${val}"`;
      case "number":
        return `${key} = 1`;
      case "array":
        if (
          task_schema.properties[key].items.type === "object" &&
          task_schema.properties[key].items.properties.id
        ) {
          return `${key} {
						id = "foo"
						name = "bar"
					}`;
        }
		if (task_schema.properties[key].items.example) {
			return `${key} = ["${task_schema.properties[key].items.example}"]`;
		} else {
			return `${key} = ["foo"]`;
		}
      case "object":
        return `${key} = {
					id = "foo"
					name = "bar"
				}`;
    }
  });
}
