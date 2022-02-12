package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

var resourceWorkloadSetSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"workload_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"set_id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"workload_resource": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"service_vendor": {
					Type:     schema.TypeString,
					Required: true,
				},
				"resource_id": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	},
}

func resourceWorkloadSet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkloadSetCreate,
		ReadContext:   resourceWorkloadSetRead,
		UpdateContext: resourceWorkloadSetUpdate,
		DeleteContext: resourceWorkloadSetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceWorkloadSetSchema,
	}
}

func workloadSetFromResourceData(d *schema.ResourceData) (string, string, []WorkloadSetItem) {
	var workloadSet []WorkloadSetItem
	workloadId, setId := d.Get("workload_id").(string), ""
	if d.Id() != "" {
		ids := strings.Split(d.Id(), "/")
		workloadId, setId = ids[0], ids[1]
	}

	items := d.Get("workload_resource").(interface{})
	for _, raw := range items.(*schema.Set).List() {
		mapRaw := raw.(map[string]interface{})
		obj := WorkloadSetItem{
			ResourceId:    mapRaw["resource_id"].(string),
			ServiceVendor: mapRaw["service_vendor"].(string),
		}
		workloadSet = append(workloadSet, obj)
	}
	return workloadId, setId, workloadSet
}

func workloadSetToResourceData(d *schema.ResourceData, w []WorkloadSetItem) {
	var workloadResourceList []map[string]string

	for _, v := range w {
		m := make(map[string]string)
		m["resource_id"] = v.ResourceId
		m["service_vendor"] = v.ServiceVendor
		workloadResourceList = append(workloadResourceList, m)
	}

	d.Set("workload_resource", workloadResourceList)
	ids := strings.Split(d.Id(), "/")
	workloadId := ids[0]
	setId := ids[1]
	d.Set("workload_id", workloadId)
	d.Set("set_id", setId)
}

func resourceWorkloadSetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics
	workloadId, setId, _ := workloadSetFromResourceData(d)
	workloadSet, err := c.GetWorkloadSet(workloadId, setId)
	if err != nil {
		return diag.FromErr(err)
	}
	workloadSetToResourceData(d, workloadSet.Payload)
	return diags
}

func resourceWorkloadSetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics
	workloadId, _, workloadSet := workloadSetFromResourceData(d)
	setId := getUniqueId()
	_, err := c.CreateWorkloadSet(workloadId, setId, workloadSet)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(workloadId + "/" + setId)
	resourceWorkloadSetRead(ctx, d, m)
	return diags
}

func resourceWorkloadSetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	workloadId, setId, workloadSet := workloadSetFromResourceData(d)
	c := m.(Client)
	_, err := c.UpdateWorkloadSet(workloadId, setId, workloadSet)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceWorkloadSetRead(ctx, d, m)
}

func resourceWorkloadSetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)
	var diags diag.Diagnostics

	workloadId, setId, _ := workloadSetFromResourceData(d)
	err := c.DeleteWorkloadSet(workloadId, setId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
