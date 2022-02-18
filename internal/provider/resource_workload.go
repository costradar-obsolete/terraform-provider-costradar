package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var resourceWorkloadSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"owners": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"tags": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}

func resourceWorkload() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkloadCreate,
		ReadContext:   resourceWorkloadRead,
		UpdateContext: resourceWorkloadUpdate,
		DeleteContext: resourceWorkloadDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceWorkloadSchema,
	}
}

func workloadFromResourceData(d *schema.ResourceData) Workload {
	workload := Workload{
		ID:          d.Get("id").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Owners:      d.Get("owners").([]interface{}),
		Tags:        d.Get("tags").(map[string]interface{}),
	}
	return workload
}

func workloadToResourceData(d *schema.ResourceData, w Workload) {
	d.Set("name", w.Name)
	if nilIfEmpty(w.Owners) != nil {
		d.Set("owners", w.Owners)
	}
	if nilIfEmpty(w.Tags) != nil {
		d.Set("tags", w.Tags)
	}
	if nilIfEmpty(w.Description) != nil {
		d.Set("description", w.Description)
	}
}

func resourceWorkloadRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	id := d.Id()
	workload, err := c.GetWorkload(id)
	if err != nil {
		return diag.FromErr(err)
	}
	workloadToResourceData(d, workload.Payload)
	return diags
}

func resourceWorkloadCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	var workload = workloadFromResourceData(d)

	w, err := c.CreateWorkload(workload)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceWorkloadRead(ctx, d, m)
	d.SetId(w.Payload.ID)
	return diags
}

func resourceWorkloadUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	workload := workloadFromResourceData(d)
	c := m.(Client)
	_, err := c.UpdateWorkload(workload)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceWorkloadRead(ctx, d, m)
}

func resourceWorkloadDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)
	var diags diag.Diagnostics

	workloadId := d.Id()

	err := c.DeleteWorkload(workloadId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
