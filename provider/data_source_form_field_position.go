package provider

import (
	"context"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rootlyhq/terraform-provider-rootly/client"
	rootlygo "github.com/rootlyhq/terraform-provider-rootly/schema"
)

func dataSourceFormFieldPosition() *schema.Resource {
	return &schema.Resource {
		ReadContext: dataSourceFormFieldPositionRead,
		Schema: map[string]*schema.Schema {
			"id": &schema.Schema {
				Type: schema.TypeString,
				Computed: true,
			},
			
			"form": &schema.Schema {
				Type: schema.TypeString,
				Computed: true,
				Optional: true,
			},
			
		},
	}
}

func dataSourceFormFieldPositionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	params := new(rootlygo.ListFormFieldPositionsParams)
	page_size := 1
	params.PageSize = &page_size

	
				if value, ok := d.GetOkExists("form"); ok {
					form := value.(string)
					params.FilterForm = &form
				}
			

	form_field_id := d.Get("form_field_id").(string)
			items, err := c.ListFormFieldPositions(form_field_id, params)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(items) == 0 {
		return diag.Errorf("form_field_position not found")
	}
	item, _ := items[0].(*client.FormFieldPosition)

	d.SetId(item.ID)

	return nil
}
