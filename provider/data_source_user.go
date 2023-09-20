package provider

import (
	"context"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rootlyhq/terraform-provider-rootly/client"
	rootlygo "github.com/rootlyhq/terraform-provider-rootly/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource {
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema {
			"id": &schema.Schema {
				Type: schema.TypeString,
				Computed: true,
			},
			
			"email": &schema.Schema {
				Type: schema.TypeString,
				Computed: true,
				Optional: true,
			},
			

				"created_at": &schema.Schema {
					Type: schema.TypeMap,
					Description: "Filter by date range using 'lt' and 'gt'.",
					Optional: true,
				},
				
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	params := new(rootlygo.ListUsersParams)
	page_size := 1
	params.PageSize = &page_size

	
				if value, ok := d.GetOkExists("email"); ok {
					email := value.(string)
					params.FilterEmail = &email
				}
			

				created_at_gt := d.Get("created_at").(map[string]interface{})
				if value, exists := created_at_gt["gt"]; exists {
					v := value.(string)
					params.FilterCreatedAtGt = &v
				}
			

				created_at_lt := d.Get("created_at").(map[string]interface{})
				if value, exists := created_at_lt["lt"]; exists {
					v := value.(string)
					params.FilterCreatedAtLt = &v
				}
			

	items, err := c.ListUsers(params)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(items) == 0 {
		return diag.Errorf("user not found")
	}
	item, _ := items[0].(*client.User)

	d.SetId(item.ID)

	return nil
}
