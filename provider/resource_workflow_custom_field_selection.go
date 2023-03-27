package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rootlyhq/terraform-provider-rootly/client"
)

func resourceWorkflowCustomFieldSelection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkflowCustomFieldSelectionCreate,
		ReadContext:   resourceWorkflowCustomFieldSelectionRead,
		UpdateContext: resourceWorkflowCustomFieldSelectionUpdate,
		DeleteContext: resourceWorkflowCustomFieldSelectionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			"workflow_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    true,
				Description: "The workflow for this selection",
			},

			"custom_field_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    false,
				Required:    true,
				Optional:    false,
				ForceNew:    false,
				Description: "The custom field for this selection",
			},

			"incident_condition": &schema.Schema{
				Type:        schema.TypeString,
				Default:     "ANY",
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The trigger condition. Value must be one of `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `NONE`, `SET`, `UNSET`.",
			},

			"values": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "",
			},

			"selected_option_ids": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceWorkflowCustomFieldSelectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	tflog.Trace(ctx, fmt.Sprintf("Creating WorkflowCustomFieldSelection"))

	s := &client.WorkflowCustomFieldSelection{}

	if value, ok := d.GetOkExists("workflow_id"); ok {
		s.WorkflowId = value.(string)
	}
	if value, ok := d.GetOkExists("custom_field_id"); ok {
		s.CustomFieldId = value.(int)
	}
	if value, ok := d.GetOkExists("incident_condition"); ok {
		s.IncidentCondition = value.(string)
	}
	if value, ok := d.GetOkExists("values"); ok {
		s.Values = value.([]interface{})
	}
	if value, ok := d.GetOkExists("selected_option_ids"); ok {
		s.SelectedOptionIds = value.([]interface{})
	}

	res, err := c.CreateWorkflowCustomFieldSelection(s)
	if err != nil {
		return diag.Errorf("Error creating workflow_custom_field_selection: %s", err.Error())
	}

	d.SetId(res.ID)
	tflog.Trace(ctx, fmt.Sprintf("created a workflow_custom_field_selection resource: %s", d.Id()))

	return resourceWorkflowCustomFieldSelectionRead(ctx, d, meta)
}

func resourceWorkflowCustomFieldSelectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading WorkflowCustomFieldSelection: %s", d.Id()))

	item, err := c.GetWorkflowCustomFieldSelection(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowCustomFieldSelection (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error reading workflow_custom_field_selection: %s", d.Id())
	}

	d.Set("workflow_id", item.WorkflowId)
	d.Set("custom_field_id", item.CustomFieldId)
	d.Set("incident_condition", item.IncidentCondition)
	d.Set("values", item.Values)
	d.Set("selected_option_ids", item.SelectedOptionIds)

	return nil
}

func resourceWorkflowCustomFieldSelectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating WorkflowCustomFieldSelection: %s", d.Id()))

	s := &client.WorkflowCustomFieldSelection{}

	if d.HasChange("workflow_id") {
		s.WorkflowId = d.Get("workflow_id").(string)
	}
	if d.HasChange("custom_field_id") {
		s.CustomFieldId = d.Get("custom_field_id").(int)
	}
	if d.HasChange("incident_condition") {
		s.IncidentCondition = d.Get("incident_condition").(string)
	}
	if d.HasChange("values") {
		s.Values = d.Get("values").([]interface{})
	}
	if d.HasChange("selected_option_ids") {
		s.SelectedOptionIds = d.Get("selected_option_ids").([]interface{})
	}

	_, err := c.UpdateWorkflowCustomFieldSelection(d.Id(), s)
	if err != nil {
		return diag.Errorf("Error updating workflow_custom_field_selection: %s", err.Error())
	}

	return resourceWorkflowCustomFieldSelectionRead(ctx, d, meta)
}

func resourceWorkflowCustomFieldSelectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting WorkflowCustomFieldSelection: %s", d.Id()))

	err := c.DeleteWorkflowCustomFieldSelection(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowCustomFieldSelection (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting workflow_custom_field_selection: %s", err.Error())
	}

	d.SetId("")

	return nil
}
