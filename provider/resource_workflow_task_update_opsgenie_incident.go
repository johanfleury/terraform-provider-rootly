package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/rootlyhq/terraform-provider-rootly/client"
)

func resourceWorkflowTaskUpdateOpsgenieIncident() *schema.Resource {
	return &schema.Resource{
		Description: "Manages workflow update_opsgenie_incident task.",

		CreateContext: resourceWorkflowTaskUpdateOpsgenieIncidentCreate,
		ReadContext:   resourceWorkflowTaskUpdateOpsgenieIncidentRead,
		UpdateContext: resourceWorkflowTaskUpdateOpsgenieIncidentUpdate,
		DeleteContext: resourceWorkflowTaskUpdateOpsgenieIncidentDelete,
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
							Default:  "update_opsgenie_incident",
							ValidateFunc: validation.StringInSlice([]string{
								"update_opsgenie_incident",
							}, false),
						},
						"opsgenie_incident_id": &schema.Schema{
							Description: "The Opsgenie incident ID, this can also be a Rootly incident variable ",
							Type:        schema.TypeString,
							Required:    true,
						},
						"message": &schema.Schema{
							Description: "Message of the alert",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"description": &schema.Schema{
							Description: "Description field of the alert that is generally used to provide a detailed information about the alert",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"status": &schema.Schema{
							Description: "Value must be one of `resolve`, `open`, `auto`.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     nil,
							ValidateFunc: validation.StringInSlice([]string{
								"resolve",
								"open",
								"auto",
							}, false),
						},
						"priority": &schema.Schema{
							Description: "Value must be one of `P1`, `P2`, `P3`, `P4`, `auto`.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     nil,
							ValidateFunc: validation.StringInSlice([]string{
								"P1",
								"P2",
								"P3",
								"P4",
								"auto",
							}, false),
						},
					},
				},
			},
		},
	}
}

func resourceWorkflowTaskUpdateOpsgenieIncidentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	return resourceWorkflowTaskUpdateOpsgenieIncidentRead(ctx, d, meta)
}

func resourceWorkflowTaskUpdateOpsgenieIncidentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading workflow task: %s", d.Id()))

	res, err := c.GetWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskUpdateOpsgenieIncident (%s) not found, removing from state", d.Id()))
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

func resourceWorkflowTaskUpdateOpsgenieIncidentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	return resourceWorkflowTaskUpdateOpsgenieIncidentRead(ctx, d, meta)
}

func resourceWorkflowTaskUpdateOpsgenieIncidentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting workflow task: %s", d.Id()))

	err := c.DeleteWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskUpdateOpsgenieIncident (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting workflow task: %s", err.Error())
	}

	d.SetId("")

	return nil
}
