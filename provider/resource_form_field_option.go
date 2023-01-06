package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rootlyhq/terraform-provider-rootly/client"
	"github.com/rootlyhq/terraform-provider-rootly/tools"
)

func resourceFormFieldOption() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFormFieldOptionCreate,
		ReadContext:   resourceFormFieldOptionRead,
		UpdateContext: resourceFormFieldOptionUpdate,
		DeleteContext: resourceFormFieldOptionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			"form_field_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    false,
				Required:    true,
				Optional:    false,
				ForceNew:    true,
				Description: "The ID of the parent custom field",
			},

			"value": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    false,
				Required:    true,
				Optional:    false,
				ForceNew:    false,
				Description: "The value of the form_field_option",
			},

			"color": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The hex color of the form_field_option",
			},

			"default": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "",
			},

			"position": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The position of the form_field_option",
			},
		},
	}
}

func resourceFormFieldOptionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	tflog.Trace(ctx, fmt.Sprintf("Creating FormFieldOption"))

	s := &client.FormFieldOption{}

	if value, ok := d.GetOkExists("form_field_id"); ok {
		s.FormFieldId = value.(string)
	}
	if value, ok := d.GetOkExists("value"); ok {
		s.Value = value.(string)
	}
	if value, ok := d.GetOkExists("color"); ok {
		s.Color = value.(string)
	}
	if value, ok := d.GetOkExists("default"); ok {
		s.Default = tools.Bool(value.(bool))
	}
	if value, ok := d.GetOkExists("position"); ok {
		s.Position = value.(int)
	}

	res, err := c.CreateFormFieldOption(s)
	if err != nil {
		return diag.Errorf("Error creating form_field_option: %s", err.Error())
	}

	d.SetId(res.ID)
	tflog.Trace(ctx, fmt.Sprintf("created a form_field_option resource: %s", d.Id()))

	return resourceFormFieldOptionRead(ctx, d, meta)
}

func resourceFormFieldOptionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading FormFieldOption: %s", d.Id()))

	item, err := c.GetFormFieldOption(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("FormFieldOption (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error reading form_field_option: %s", d.Id())
	}

	d.Set("form_field_id", item.FormFieldId)
	d.Set("value", item.Value)
	d.Set("color", item.Color)
	d.Set("default", item.Default)
	d.Set("position", item.Position)

	return nil
}

func resourceFormFieldOptionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating FormFieldOption: %s", d.Id()))

	s := &client.FormFieldOption{}

	if d.HasChange("form_field_id") {
		s.FormFieldId = d.Get("form_field_id").(string)
	}
	if d.HasChange("value") {
		s.Value = d.Get("value").(string)
	}
	if d.HasChange("color") {
		s.Color = d.Get("color").(string)
	}
	if d.HasChange("default") {
		s.Default = tools.Bool(d.Get("default").(bool))
	}
	if d.HasChange("position") {
		s.Position = d.Get("position").(int)
	}

	_, err := c.UpdateFormFieldOption(d.Id(), s)
	if err != nil {
		return diag.Errorf("Error updating form_field_option: %s", err.Error())
	}

	return resourceFormFieldOptionRead(ctx, d, meta)
}

func resourceFormFieldOptionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting FormFieldOption: %s", d.Id()))

	err := c.DeleteFormFieldOption(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("FormFieldOption (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting form_field_option: %s", err.Error())
	}

	d.SetId("")

	return nil
}
