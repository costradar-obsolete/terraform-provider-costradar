package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var resourceTenantSchema = map[string]*schema.Schema{
	"alias": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"icon_url": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"auth": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"client_id": {
					Type:     schema.TypeString,
					Required: true,
				},
				"client_secret": {
					Type:     schema.TypeString,
					Required: true,
				},
				"server_metadata_url": {
					Type:     schema.TypeString,
					Required: true,
				},
				"client_kwargs": {
					Type:     schema.TypeMap,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"email_domains": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	},
}

func resourceTenant() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTenantCreate,
		ReadContext:   resourceTenantRead,
		UpdateContext: resourceTenantUpdate,
		DeleteContext: resourceTenantDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceTenantSchema,
	}
}

func tenantFromResourceData(d *schema.ResourceData) Tenant {
	var auth TenantAuth
	authData := d.Get("auth.0").(map[string]interface{})

	if v, ok := authData["client_id"].(string); ok {
		auth.ClientId = v
	}
	if v, ok := authData["client_secret"].(string); ok {
		auth.ClientSecret = v
	}
	if v, ok := authData["server_metadata_url"].(string); ok {
		auth.ServerMetadataUrl = v
	}
	if v, ok := authData["client_kwargs"].(map[string]interface{}); ok {
		auth.ClientKwargs = v
	}
	if v, ok := authData["email_domains"].([]interface{}); ok {
		auth.EmailDomains = v
	}
	tenant := Tenant{
		Alias:       d.Get("alias").(string),
		Description: d.Get("description").(string),
		IconUrl:     d.Get("icon_url").(string),
		Auth:        auth,
	}
	return tenant
}

func tenantToResourceData(d *schema.ResourceData, t Tenant) {

	var authList []map[string]interface{}
	auth := make(map[string]interface{})
	auth["client_id"] = t.Auth.ClientId
	auth["client_secret"] = t.Auth.ClientSecret
	auth["server_metadata_url"] = t.Auth.ServerMetadataUrl
	auth["client_kwargs"] = t.Auth.ClientKwargs.(map[string]string)
	auth["email_domains"] = t.Auth.EmailDomains
	authList = append(authList, auth)
	d.Set("alias", t.Alias)
	d.Set("description", t.Description)
	d.Set("icon_url", t.IconUrl)
	d.Set("auth", authList)
}

func resourceTenantRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	tenant, err := c.GetTenant()
	if err != nil {
		return diag.FromErr(err)
	}
	tenantToResourceData(d, tenant.Payload)
	return diags
}

func resourceTenantCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	tenant := tenantFromResourceData(d)

	_, err := c.UpdateTenant(tenant)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(getUniqueId())
	resourceTenantRead(ctx, d, m)
	return diags
}

func resourceTenantUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tenant := tenantFromResourceData(d)
	c := m.(Client)
	_, err := c.UpdateTenant(tenant)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceTenantRead(ctx, d, m)
}

func resourceTenantDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.SetId("")
	return diags
}
