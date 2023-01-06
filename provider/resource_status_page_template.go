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

func resourceStatusPageTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceStatusPageTemplateCreate,
		ReadContext:   resourceStatusPageTemplateRead,
		UpdateContext: resourceStatusPageTemplateUpdate,
		DeleteContext: resourceStatusPageTemplateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			"status_page_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    true,
				Description: "",
			},

			"title": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    false,
				Required:    true,
				Optional:    false,
				ForceNew:    false,
				Description: "Title of the template",
			},

			"body": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    false,
				Required:    true,
				Optional:    false,
				ForceNew:    false,
				Description: "Description of the event the template will populate",
			},

			"update_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "Status of the event the template will populate",
			},

			"should_notify_subscribers": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "Controls if incident subscribers should be notified",
			},

			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},

			"position": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "Position of the workflow task",
			},
		},
	}
}

func resourceStatusPageTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	tflog.Trace(ctx, fmt.Sprintf("Creating StatusPageTemplate"))

	s := &client.StatusPageTemplate{}

	if value, ok := d.GetOkExists("status_page_id"); ok {
		s.StatusPageId = value.(string)
	}
	if value, ok := d.GetOkExists("title"); ok {
		s.Title = value.(string)
	}
	if value, ok := d.GetOkExists("body"); ok {
		s.Body = value.(string)
	}
	if value, ok := d.GetOkExists("update_status"); ok {
		s.UpdateStatus = value.(string)
	}
	if value, ok := d.GetOkExists("should_notify_subscribers"); ok {
		s.ShouldNotifySubscribers = tools.Bool(value.(bool))
	}
	if value, ok := d.GetOkExists("enabled"); ok {
		s.Enabled = tools.Bool(value.(bool))
	}
	if value, ok := d.GetOkExists("position"); ok {
		s.Position = value.(int)
	}

	res, err := c.CreateStatusPageTemplate(s)
	if err != nil {
		return diag.Errorf("Error creating status_page_template: %s", err.Error())
	}

	d.SetId(res.ID)
	tflog.Trace(ctx, fmt.Sprintf("created a status_page_template resource: %s", d.Id()))

	return resourceStatusPageTemplateRead(ctx, d, meta)
}

func resourceStatusPageTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading StatusPageTemplate: %s", d.Id()))

	item, err := c.GetStatusPageTemplate(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("StatusPageTemplate (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error reading status_page_template: %s", d.Id())
	}

	d.Set("status_page_id", item.StatusPageId)
	d.Set("title", item.Title)
	d.Set("body", item.Body)
	d.Set("update_status", item.UpdateStatus)
	d.Set("should_notify_subscribers", item.ShouldNotifySubscribers)
	d.Set("enabled", item.Enabled)
	d.Set("position", item.Position)

	return nil
}

func resourceStatusPageTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating StatusPageTemplate: %s", d.Id()))

	s := &client.StatusPageTemplate{}

	if d.HasChange("status_page_id") {
		s.StatusPageId = d.Get("status_page_id").(string)
	}
	if d.HasChange("title") {
		s.Title = d.Get("title").(string)
	}
	if d.HasChange("body") {
		s.Body = d.Get("body").(string)
	}
	if d.HasChange("update_status") {
		s.UpdateStatus = d.Get("update_status").(string)
	}
	if d.HasChange("should_notify_subscribers") {
		s.ShouldNotifySubscribers = tools.Bool(d.Get("should_notify_subscribers").(bool))
	}
	if d.HasChange("enabled") {
		s.Enabled = tools.Bool(d.Get("enabled").(bool))
	}
	if d.HasChange("position") {
		s.Position = d.Get("position").(int)
	}

	_, err := c.UpdateStatusPageTemplate(d.Id(), s)
	if err != nil {
		return diag.Errorf("Error updating status_page_template: %s", err.Error())
	}

	return resourceStatusPageTemplateRead(ctx, d, meta)
}

func resourceStatusPageTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting StatusPageTemplate: %s", d.Id()))

	err := c.DeleteStatusPageTemplate(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("StatusPageTemplate (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting status_page_template: %s", err.Error())
	}

	d.SetId("")

	return nil
}
