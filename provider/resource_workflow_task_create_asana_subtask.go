package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"context"
	"fmt"

	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/rootlyhq/terraform-provider-rootly/client"
)

func resourceWorkflowTaskCreateAsanaSubtask() *schema.Resource {
	return &schema.Resource{
		Description: "Manages workflow create_asana_subtask task.",

		CreateContext: resourceWorkflowTaskCreateAsanaSubtaskCreate,
		ReadContext:   resourceWorkflowTaskCreateAsanaSubtaskRead,
		UpdateContext: resourceWorkflowTaskCreateAsanaSubtaskUpdate,
		DeleteContext: resourceWorkflowTaskCreateAsanaSubtaskDelete,
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
							Default:  "create_asana_subtask",
							ValidateFunc: validation.StringInSlice([]string{
								"create_asana_subtask",
							}, false),
						},
						"parent_task_id": &schema.Schema{
							Description: "The parent task id",
							Type:        schema.TypeString,
							Required:    true,
						},
						"title": &schema.Schema{
							Description: "The subtask title",
							Type:        schema.TypeString,
							Required:    true,
						},
						"notes": &schema.Schema{
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"assign_user_email": &schema.Schema{
							Description: "The assigned user's email.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"completion": &schema.Schema{
							Description: "Map must contain two fields, `id` and `name`. ",
							Type:        schema.TypeMap,
							Required:    true,
						},
						"custom_fields_mapping": &schema.Schema{
							Description: "Custom field mappings. Can contain liquid markup and need to be valid JSON.",
							Type:        schema.TypeString,
							Optional:    true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								t := &testing.T{}
								assert := assert.New(t)
								return assert.JSONEq(old, new)
							},
							Default: "{}",
						},
						"dependency_direction": &schema.Schema{
							Description: "Value must be one of `blocking`, `blocked_by`.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "blocking",
							ValidateFunc: validation.StringInSlice([]string{
								"blocking",
								"blocked_by",
							}, false),
						},
						"dependent_task_ids": &schema.Schema{
							Description: "Dependent task ids. Supports liquid syntax.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceWorkflowTaskCreateAsanaSubtaskCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	workflowId := d.Get("workflow_id").(string)
	position := d.Get("position").(int)
	skipOnFailure := d.Get("skip_on_failure").(bool)
	enabled := d.Get("enabled").(bool)
	taskParams := d.Get("task_params").([]interface{})[0].(map[string]interface{})

	tflog.Trace(ctx, fmt.Sprintf("Creating workflow task: %s", workflowId))

	s := &client.WorkflowTask{
		WorkflowId:    workflowId,
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

	return resourceWorkflowTaskCreateAsanaSubtaskRead(ctx, d, meta)
}

func resourceWorkflowTaskCreateAsanaSubtaskRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading workflow task: %s", d.Id()))

	res, err := c.GetWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskCreateAsanaSubtask (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error reading workflow task: %s", d.Id())
	}

	d.Set("workflow_id", res.WorkflowId)
	d.Set("position", res.Position)
	d.Set("skip_on_failure", res.SkipOnFailure)
	d.Set("enabled", res.Enabled)
	tps := make([]interface{}, 1, 1)
	tps[0] = res.TaskParams
	d.Set("task_params", tps)

	return nil
}

func resourceWorkflowTaskCreateAsanaSubtaskUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating workflow task: %s", d.Id()))

	workflowId := d.Get("workflow_id").(string)
	position := d.Get("position").(int)
	skipOnFailure := d.Get("skip_on_failure").(bool)
	enabled := d.Get("enabled").(bool)
	taskParams := d.Get("task_params").([]interface{})[0].(map[string]interface{})

	s := &client.WorkflowTask{
		WorkflowId:    workflowId,
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

	return resourceWorkflowTaskCreateAsanaSubtaskRead(ctx, d, meta)
}

func resourceWorkflowTaskCreateAsanaSubtaskDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting workflow task: %s", d.Id()))

	err := c.DeleteWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskCreateAsanaSubtask (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting workflow task: %s", err.Error())
	}

	d.SetId("")

	return nil
}
