package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rootlyhq/terraform-provider-rootly/client"
)

func resourceStatusPage() *schema.Resource{
	return &schema.Resource{
		CreateContext: resourceStatusPageCreate,
		ReadContext: resourceStatusPageRead,
		UpdateContext: resourceStatusPageUpdate,
		DeleteContext: resourceStatusPageDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			
			"title": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
				Required: false,
				Optional: true,
				Description: "The title of the status page",
			},
			

			"description": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
				Required: false,
				Optional: true,
				Description: "The description of the status page",
			},
			

			"header_color": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
				Required: false,
				Optional: true,
				Description: "The color of the header. Eg. \"#0061F2\"",
			},
			

			"footer_color": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
				Required: false,
				Optional: true,
				Description: "The color of the footer. Eg. \"#1F2F41\"",
			},
			

			"allow_search_engine_index": &schema.Schema{
				Type: schema.TypeBool,
				Computed: true,
				Required: false,
				Optional: true,
				Description: "Allow search engines to include your public status page in search results.",
			},
			

			"show_uptime": &schema.Schema{
				Type: schema.TypeBool,
				Computed: true,
				Required: false,
				Optional: true,
				Description: "Show uptime",
			},
			

			"show_uptime_last_days": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
				Required: false,
				Optional: true,
				Description: "Show uptime over x days",
			},
			

			"public": &schema.Schema{
				Type: schema.TypeBool,
				Computed: true,
				Required: false,
				Optional: true,
				Description: "Make the status page accessible to the public",
			},
			

			"enabled": &schema.Schema{
				Type: schema.TypeBool,
				Computed: true,
				Required: false,
				Optional: true,
				Description: "Enabled / Disable the status page",
			},
			
		},
	}
}

func resourceStatusPageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	tflog.Trace(ctx, fmt.Sprintf("Creating StatusPage"))

	s := &client.StatusPage{}

	  if value, ok := d.GetOkExists("title"); ok {
		s.Title = value.(string)
	}
    if value, ok := d.GetOkExists("description"); ok {
		s.Description = value.(string)
	}
    if value, ok := d.GetOkExists("header_color"); ok {
		s.HeaderColor = value.(string)
	}
    if value, ok := d.GetOkExists("footer_color"); ok {
		s.FooterColor = value.(string)
	}
    if value, ok := d.GetOkExists("allow_search_engine_index"); ok {
		s.AllowSearchEngineIndex = value.(bool)
	}
    if value, ok := d.GetOkExists("show_uptime"); ok {
		s.ShowUptime = value.(bool)
	}
    if value, ok := d.GetOkExists("show_uptime_last_days"); ok {
		s.ShowUptimeLastDays = value.(string)
	}
    if value, ok := d.GetOkExists("public"); ok {
		s.Public = value.(bool)
	}
    if value, ok := d.GetOkExists("enabled"); ok {
		s.Enabled = value.(bool)
	}

	res, err := c.CreateStatusPage(s)
	if err != nil {
		return diag.Errorf("Error creating status_page: %s", err.Error())
	}

	d.SetId(res.ID)
	tflog.Trace(ctx, fmt.Sprintf("created a status_page resource: %s", d.Id()))

	return resourceStatusPageRead(ctx, d, meta)
}

func resourceStatusPageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading StatusPage: %s", d.Id()))

	item, err := c.GetStatusPage(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("StatusPage (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error reading status_page: %s", d.Id())
	}

	d.Set("title", item.Title)
  d.Set("description", item.Description)
  d.Set("header_color", item.HeaderColor)
  d.Set("footer_color", item.FooterColor)
  d.Set("allow_search_engine_index", item.AllowSearchEngineIndex)
  d.Set("show_uptime", item.ShowUptime)
  d.Set("show_uptime_last_days", item.ShowUptimeLastDays)
  d.Set("public", item.Public)
  d.Set("enabled", item.Enabled)

	return nil
}

func resourceStatusPageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating StatusPage: %s", d.Id()))

	s := &client.StatusPage{}

	  if d.HasChange("title") {
		s.Title = d.Get("title").(string)
	}
    if d.HasChange("description") {
		s.Description = d.Get("description").(string)
	}
    if d.HasChange("header_color") {
		s.HeaderColor = d.Get("header_color").(string)
	}
    if d.HasChange("footer_color") {
		s.FooterColor = d.Get("footer_color").(string)
	}
    if d.HasChange("allow_search_engine_index") {
		s.AllowSearchEngineIndex = d.Get("allow_search_engine_index").(bool)
	}
    if d.HasChange("show_uptime") {
		s.ShowUptime = d.Get("show_uptime").(bool)
	}
    if d.HasChange("show_uptime_last_days") {
		s.ShowUptimeLastDays = d.Get("show_uptime_last_days").(string)
	}
    if d.HasChange("public") {
		s.Public = d.Get("public").(bool)
	}
    if d.HasChange("enabled") {
		s.Enabled = d.Get("enabled").(bool)
	}

	_, err := c.UpdateStatusPage(d.Id(), s)
	if err != nil {
		return diag.Errorf("Error updating status_page: %s", err.Error())
	}

	return resourceStatusPageRead(ctx, d, meta)
}

func resourceStatusPageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting StatusPage: %s", d.Id()))

	err := c.DeleteStatusPage(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("StatusPage (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting status_page: %s", err.Error())
	}

	d.SetId("")

	return nil
}
