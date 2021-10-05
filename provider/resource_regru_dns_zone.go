package provider

import (
	"context"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func zone() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		CreateContext: zoneContexCreate,
		ReadContext:   zoneContexRead,
		Update:        zoneAction,
		Delete:        zoneAction,
	}
}

func zoneContexRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func zoneAction(data *schema.ResourceData, i interface{}) error {
	return nil
}

func zoneContexCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	data.SetId(uuid.New().String())
	return diags
}
