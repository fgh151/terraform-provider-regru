package provider

import (
	"context"
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
		ReadContext:   resourceRecordRead,
		DeleteContext: resourceDeleteRecord,
		UpdateContext: resourceUpdateContext,
	}
}

func resourceUpdateContext(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(RegruProvider)
	var diags diag.Diagnostics

	ar, err, resp := c.GetRecords(data.Get("zone").(string))

	var record DnsRecord
	record_type := data.Get("type").(string)
	value := data.Get("value").(string)
	for _, r := range ar {
		if r.Type == record_type && r.Value == value {
			record = r
		}
	}

	err, resp = c.DeleteRecord(record)
	err, body := c.AddRecord(record)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Summary:  "read",
			Severity: diag.Error,
			Detail:   err.Error() + fmt.Sprintf("%V", ar) + string(resp) + string(body),
		})
	}

	return diags
}

func resourceCreateRecord(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(RegruProvider)

	var diags diag.Diagnostics

	err, body := c.AddRecord(DnsRecord{
		Subdomain: data.Get("host").(string),
		Host:      data.Get("host").(string),
		Type:      data.Get("type").(string),
		Value:     data.Get("value").(string),
		Ttl:       data.Get("ttl").(int),
		Domain:    data.Get("zone").(string),
	})

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Create record error",
			Detail:   string(body),
		})
		return diags
	}

	data.SetId(uuid.New().String())

	return nil
}

func resourceRecordRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(RegruProvider)
	var diags diag.Diagnostics

	ar, err, resp := c.GetRecords(data.Get("zone").(string))

	var record DnsRecord
	record_type := data.Get("type").(string)
	value := data.Get("value").(string)
	for _, r := range ar {
		if r.Type == record_type && r.Value == value {
			record = r
		}
	}

	if record != (DnsRecord{}) {
		err = data.Set("host", record.Host)
		err = data.Set("type", record.Type)
		err = data.Set("value", record.Value)
		err = data.Set("ttl", record.Ttl)
		err = data.Set("zone", record.Domain)
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

func resourceDeleteRecord(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {

	c := i.(RegruProvider)
	var diags diag.Diagnostics

	err, resp := c.DeleteRecord(DnsRecord{
		Host:  data.Get("host").(string),
		Value: data.Get("value").(string),
		Type:  data.Get("type").(string),
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Summary:  "read",
			Severity: diag.Error,
			Detail:   err.Error() + string(resp),
		})
	}

	return diags
}
