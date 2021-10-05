package provider

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func record() *schema.Resource {
	return &schema.Resource{

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

		CreateContext: resourceCreateRecord,

		ReadContext: dataSourceRecordRead,

		//Read:   resourceReadRecord,
		Update: resourceUpdateRecord,
		Delete: resourceDeleteRecord,
	}
}

func resourceCreateRecord(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(RegruProvider)

	var diags diag.Diagnostics

	_, body, req := c.AddRecord(DnsRecord{
		Subdomain: d.Get("host").(string),
		Host:      d.Get("host").(string),
		Type:      d.Get("type").(string),
		Value:     d.Get("value").(string),
		Ttl:       d.Get("ttl").(int),

		Domain: "", //TODO get from parent resource
	})

	str := fmt.Sprintf("Underlying Value: %v\n", c)

	//if err != nil {
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Create record error",
		Detail:   str + " " + string(body) + " " + string(req),
	})
	//}

	return diags
}

func dataSourceRecordRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(RegruProvider)
	var diags diag.Diagnostics
	str := fmt.Sprintf("Underlying Value: %v\n", c)

	diags = append(diags, diag.Diagnostic{
		Summary:  "read",
		Severity: diag.Error,
		Detail:   str,
	})

	return diags
}

//func resourceCreateRecord(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
//c := i.(*RegruProvider)
//var diags diag.Diagnostics

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

//return diags
//}

func resourceDeleteRecord(data *schema.ResourceData, i interface{}) error {

	err := errors.New("Delete zone err")

	if err != nil {
		return err
	}

	return nil
}

func resourceUpdateRecord(data *schema.ResourceData, i interface{}) error {
	err := errors.New("Update zone err")

	if err != nil {
		return err
	}

	return nil
}

func resourceReadRecord(data *schema.ResourceData, i interface{}) error {
	err := errors.New("Read zone err")

	if err != nil {
		return err
	}

	return nil
}
