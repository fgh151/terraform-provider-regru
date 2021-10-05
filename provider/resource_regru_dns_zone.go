package provider

import (
	"context"
	"errors"
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
			"records": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Host",
							//ForceNew:     true,
							//ValidateFunc: validateName,
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Record type",
							//ForceNew:     true,
							//ValidateFunc: validateName,
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Record value",
							//ForceNew:     true,
							//ValidateFunc: validateName,
						},
						"ttl": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "TTL",
							//ForceNew:     true,
							//ValidateFunc: validateName,
						},
						"external_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "External ID",
							//ForceNew:     true,
							//ValidateFunc: validateName,
						},
						"additional_info": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Some info",
							//ForceNew:     true,
							//ValidateFunc: validateName,
						},
					},
				},
			},
		},

		//CreateContext: resourceCreateRecord,

		ReadContext: dataSourceZoneRead,

		//Read:   resourceReadRecord,
		Update: resourceUpdateZone,
		Delete: resourceDeleteZone,
	}
}

func dataSourceZoneRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(*RegruProvider)
	var diags diag.Diagnostics

	c.GetRecords("iqmenu.online")

	return diags
}

func resourceCreateZone(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	//c := i.(*RegruProvider)
	var diags diag.Diagnostics

	//err := c.AddRecord(DnsRecord{
	//	Host:  d.Get("host").(string),
	//	Type:  d.Get("type").(string),
	//	Value: d.Get("value").(string),
	//	Ttl:   d.Get("ttl").(int),
	//})
	//
	//if err != nil {
	//	diags = append(diags, diag.Diagnostic{
	//		Severity: diag.Error,
	//		Detail:   err.Error(),
	//	})
	//}

	return diags
}

func resourceDeleteZone(data *schema.ResourceData, i interface{}) error {

	err := errors.New("Delete zone err")

	if err != nil {
		return err
	}

	return nil
}

func resourceUpdateZone(data *schema.ResourceData, i interface{}) error {
	err := errors.New("Update zone err")

	if err != nil {
		return err
	}

	return nil
}

func resourceReadZone(data *schema.ResourceData, i interface{}) error {
	err := errors.New("Read zone err")

	if err != nil {
		return err
	}

	return nil
}
