package provider

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func record() *schema.Resource {
	return &schema.Resource{

		Schema: map[string]*schema.Schema{
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain zone",
			},
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

	err, body := c.AddRecord(DnsRecord{
		Subdomain: d.Get("host").(string),
		Host:      d.Get("host").(string),
		Type:      d.Get("type").(string),
		Value:     d.Get("value").(string),
		Ttl:       d.Get("ttl").(int),
		Domain:    d.Get("zone").(string),
	})

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Create record error",
			Detail:   string(body),
		})
		return diags
	}

	d.SetId(uuid.New().String())

	return nil
}

func dataSourceRecordRead(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(RegruProvider)
	var diags diag.Diagnostics

	ar, err, resp := c.GetRecords(d.Get("zone").(string))

	var record DnsRecord
	record_type := d.Get("type").(string)
	value := d.Get("value").(string)
	for _, r := range ar {
		if r.Type == record_type && r.Value == value {
			record = r
		}
	}

	if record != (DnsRecord{}) {
		err = d.Set("host", record.Host)
		err = d.Set("type", record.Type)
		err = d.Set("value", record.Value)
		err = d.Set("ttl", record.Ttl)
		err = d.Set("zone", record.Domain)
	}

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Summary:  "read",
			Severity: diag.Error,
			Detail:   err.Error() + fmt.Sprintf("%V", ar) + string(resp),
		})
	}

	return diags
}

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
