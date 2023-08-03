package provider

// This file was auto-generated by tools/gen_tasks.js

import (
	"context"
	"fmt"
	
	"testing"
  "github.com/stretchr/testify/assert"
	
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/rootlyhq/terraform-provider-rootly/client"
	"github.com/rootlyhq/terraform-provider-rootly/tools"
)

func resourceWorkflowTaskHttpClient() *schema.Resource {
	return &schema.Resource{
		Description: "Manages workflow http_client task.",

		CreateContext: resourceWorkflowTaskHttpClientCreate,
		ReadContext:   resourceWorkflowTaskHttpClientRead,
		UpdateContext: resourceWorkflowTaskHttpClientUpdate,
		DeleteContext: resourceWorkflowTaskHttpClientDelete,
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
							Default: "http_client",
							ValidateFunc: validation.StringInSlice([]string {
								"http_client",
							}, false),
						},
						"headers": &schema.Schema {
							Description: "JSON map of HTTP headers.",
							Type: schema.TypeString,
							Optional: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								t := &testing.T{}
								assert := assert.New(t)
								return assert.JSONEq(old, new)
							},
							Default: "{}",
						},
						"params": &schema.Schema {
							Description: "JSON map of HTTP query parameters.",
							Type: schema.TypeString,
							Optional: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								t := &testing.T{}
								assert := assert.New(t)
								return assert.JSONEq(old, new)
							},
							Default: "{}",
						},
						"body": &schema.Schema {
							Description: "HTTP body.",
							Type: schema.TypeString,
							Optional: true,
						},
						"url": &schema.Schema {
							Description: "URL endpoint.",
							Type: schema.TypeString,
							Required: true,
						},
						"event_url": &schema.Schema {
							Description: "",
							Type: schema.TypeString,
							Optional: true,
						},
						"event_message": &schema.Schema {
							Description: "",
							Type: schema.TypeString,
							Optional: true,
						},
						"method": &schema.Schema {
							Description: "HTTP method.. Value must be one of `GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`.",
							Type: schema.TypeString,
							Optional: true,
							Default: "GET",
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
"POST",
"PUT",
"DELETE",
"OPTIONS",
							}, false),
						},
						"succeed_on_status": &schema.Schema {
							Description: "HTTP status code expected. Can be a regular expression. Eg: 200, 200|203, 20[0-3]",
							Type: schema.TypeString,
							Required: true,
						},
						"post_to_incident_timeline": &schema.Schema {
							Description: "",
							Type: schema.TypeBool,
							Optional: true,
						},
						"post_to_slack_channels": &schema.Schema {
							Description: "",
							Type: schema.TypeList,
							Optional: true,
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
							},
						},
					},
				},
			},
		},
	}
}

func resourceWorkflowTaskHttpClientCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	return resourceWorkflowTaskHttpClientRead(ctx, d, meta)
}

func resourceWorkflowTaskHttpClientRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading workflow task: %s", d.Id()))

	res, err := c.GetWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskHttpClient (%s) not found, removing from state", d.Id()))
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

func resourceWorkflowTaskHttpClientUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	return resourceWorkflowTaskHttpClientRead(ctx, d, meta)
}

func resourceWorkflowTaskHttpClientDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting workflow task: %s", d.Id()))

	err := c.DeleteWorkflowTask(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("WorkflowTaskHttpClient (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting workflow task: %s", err.Error())
	}

	d.SetId("")

	return nil
}
