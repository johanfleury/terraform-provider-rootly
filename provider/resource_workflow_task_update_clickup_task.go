package provider

// This file was auto-generated by tools/generate-tasks.js

import (
	"context"
	"fmt"

	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/rootlyhq/terraform-provider-rootly/client"
	"github.com/rootlyhq/terraform-provider-rootly/tools"
)

func resourceWorkflowTaskUpdateClickupTask() *schema.Resource {
	return &schema.Resource{
		Description: "Manages workflow update_clickup_task task.",

		CreateContext: resourceWorkflowTaskUpdateClickupTaskCreate,
		ReadContext:   resourceWorkflowTaskUpdateClickupTaskRead,
		UpdateContext: resourceWorkflowTaskUpdateClickupTaskUpdate,
		DeleteContext: resourceWorkflowTaskUpdateClickupTaskDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"workflow_id": {
				Description: "The ID of the parent workflow",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Name of the workflow task",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"position": {
				Description: "The position of the workflow task (1 being top of list)",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"skip_on_failure": {
				Description: "Skip workflow task if any failures",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"enabled": {
				Description: "Enable/disable this workflow task",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"task_params": {
				Description: "The parameters for this workflow task.",
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"task_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "update_clickup_task",
							ValidateFunc: validation.StringInSlice([]string{
								"update_clickup_task",
							}, false),
						},
						"task_id": &schema.Schema{
							Description: "The task id",
							Type:        schema.TypeString,
							Required:    true,
						},
						"title": &schema.Schema{
							Description: "The task title",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"description": &schema.Schema{
							Description: "The task description",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"tags": &schema.Schema{
							Description: "The task tags",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"priority": &schema.Schema{
							Description: "Map must contain two fields, `id` and `name`. The priority id and display name",
							Type:        schema.TypeMap,
							Optional:    true,
						},
						"due_date": &schema.Schema{
							Description: "The due date",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"custom_fields_mapping": &schema.Schema{
							Description: "Custom field mappings. Can contain liquid markup and need to be valid JSON",
							Type:        schema.TypeString,
							Optional:    true,
							DiffSuppressFunc: func(k, old string, new string, d *schema.ResourceData) bool {
								var oldJSONAsInterface, newJSONAsInterface interface{}

								if err := json.Unmarshal([]byte(old), &oldJSONAsInterface); err != nil {
									return false
								}

								if err := json.Unmarshal([]byte(new), &newJSONAsInterface); err != nil {
									return false
								}

								return reflect.DeepEqual(oldJSONAsInterface, newJSONAsInterface)
							},
							Default: "{}",
						},
						"task_payload": &schema.Schema{
							Description: "Additional ClickUp task attributes. Will be merged into whatever was specified in this tasks current parameters. Can contain liquid markup and need to be valid JSON",
							Type:        schema.TypeString,
							Optional:    true,
							DiffSuppressFunc: func(k, old string, new string, d *schema.ResourceData) bool {
								var oldJSONAsInterface, newJSONAsInterface interface{}

								if err := json.Unmarshal([]byte(old), &oldJSONAsInterface); err != nil {
									return false
								}

								if err := json.Unmarshal([]byte(new), &newJSONAsInterface); err != nil {
									return false
								}

								return reflect.DeepEqual(oldJSONAsInterface, newJSONAsInterface)
							},
							Default: "{}",
						},
					},
				},
			},
		},
	}
}

func resourceWorkflowTaskUpdateClickupTaskCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	workflowId := d.Get("workflow_id").(string)
	name := d.Get("name").(string)
	position := d.Get("position").(int)
	skipOnFailure := tools.Bool(d.Get("skip_on_failure").(bool))
	enabled := tools.Bool(d.Get("enabled").(bool))
	taskParams := d.Get("task_params").([]interface{})[0].(map[string]interface{})

	tflog.Trace(ctx, fmt.Sprintf("Creating workflow task: %s", workflowId))

	s := &client.WorkflowTask{
		WorkflowId:    workflowId,
		Name:          name,
		Position:      position,
		SkipOnFailure: skipOnFailure,
		Enabled:       enabled,
		TaskParams:    taskParams,
	}

	res, err := c.CreateWorkflowTask(s)
	if err != nil {
		return diag.Errorf("Error creating workflow task: %s", err.Error())
	}

	d.SetId(res.ID)
	tflog.Trace(ctx, fmt.Sprintf("created an workflow task resource: %v (%s)", workflowId, d.Id()))

	return resourceWorkflowTaskUpdateClickupTaskRead(ctx, d, meta)
}

func resourceWorkflowTaskUpdateClickupTaskRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading workflow task: %s", d.Id()))

	res, err := c.GetWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskUpdateClickupTask (%s) not found, removing from state", d.Id()))
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

func resourceWorkflowTaskUpdateClickupTaskUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating workflow task: %s", d.Id()))

	workflowId := d.Get("workflow_id").(string)
	name := d.Get("name").(string)
	position := d.Get("position").(int)
	skipOnFailure := tools.Bool(d.Get("skip_on_failure").(bool))
	enabled := tools.Bool(d.Get("enabled").(bool))
	taskParams := d.Get("task_params").([]interface{})[0].(map[string]interface{})

	s := &client.WorkflowTask{
		WorkflowId:    workflowId,
		Name:          name,
		Position:      position,
		SkipOnFailure: skipOnFailure,
		Enabled:       enabled,
		TaskParams:    taskParams,
	}

	tflog.Debug(ctx, fmt.Sprintf("adding value: %#v", s))
	_, err := c.UpdateWorkflowTask(d.Id(), s)
	if err != nil {
		return diag.Errorf("Error updating workflow task: %s", err.Error())
	}

	return resourceWorkflowTaskUpdateClickupTaskRead(ctx, d, meta)
}

func resourceWorkflowTaskUpdateClickupTaskDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting workflow task: %s", d.Id()))

	err := c.DeleteWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskUpdateClickupTask (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting workflow task: %s", err.Error())
	}

	d.SetId("")

	return nil
}
