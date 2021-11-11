package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

import (
	"context"
)

var resourceIdentityResolverConfigSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"lambda_arn": {
		Type:     schema.TypeString,
		Required: true,
	},
	"access_config": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"reader_mode": {
					Type:     schema.TypeString,
					Required: true,
				},
				"assume_role_arn": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"assume_role_external_id": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"assume_role_session_name": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	},
}

func resourceIdentityResolverConfig() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIdentityResolverCreate,
		ReadContext:   resourceIdentityResolverRead,
		UpdateContext: resourceIdentityResolverUpdate,
		DeleteContext: resourceIdentityResolverDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceIdentityResolverConfigSchema,
	}
}

func identityResolverFromResourceData(d *schema.ResourceData) IdentityResolver {
	var accessConfig AccessConfig
	accessConfigData := d.Get("access_config.0").(map[string]interface{})
	if v, ok := accessConfigData["reader_mode"].(string); ok {
		accessConfig.ReaderMode = v
	}
	if v, ok := accessConfigData["assume_role_arn"].(string); ok {
		accessConfig.AssumeRoleArn = v
	}
	if v, ok := accessConfigData["assume_role_external_id"].(string); ok {
		accessConfig.AssumeRoleExternalId = v
	}
	if v, ok := accessConfigData["assume_role_session_name"].(string); ok {
		accessConfig.AssumeRoleSessionName = v
	}
	resolver := IdentityResolver{
		ID:           d.Get("id").(string),
		LambdaArn:    d.Get("lambda_arn").(string),
		AccessConfig: accessConfig,
	}

	return resolver
}

func identityResolverToResourceData(d *schema.ResourceData, rc IdentityResolver) {
	var accessConfigList []map[string]string
	accessConfig := make(map[string]string)
	accessConfig["reader_mode"] = rc.AccessConfig.ReaderMode
	accessConfig["assume_role_arn"] = rc.AccessConfig.AssumeRoleArn
	accessConfig["assume_role_external_id"] = rc.AccessConfig.AssumeRoleExternalId
	accessConfig["assume_role_session_name"] = rc.AccessConfig.AssumeRoleSessionName
	accessConfigList = append(accessConfigList, accessConfig)
	d.Set("lambda_arn", rc.LambdaArn)
	d.Set("access_config", accessConfigList)
}

func resourceIdentityResolverCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	var resolverConfig = identityResolverFromResourceData(d)

	s, err := c.CreateIdentityResolver(resolverConfig)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(s.Payload.ID)
	resourceIdentityResolverRead(ctx, d, m)

	return diags
}

func resourceIdentityResolverRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	resolverConfigId := d.Id()
	resolverConfig, err := c.GetIdentityResolver(resolverConfigId)
	if err != nil {
		return diag.FromErr(err)
	}
	identityResolverToResourceData(d, resolverConfig.Payload)
	return diags
}

func resourceIdentityResolverUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	resolverConfig := identityResolverFromResourceData(d)
	c := m.(Client)
	_, err := c.UpdateIdentityResolver(resolverConfig)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceIdentityResolverRead(ctx, d, m)
}

func resourceIdentityResolverDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)
	var diags diag.Diagnostics

	resolverConfigId := d.Id()

	err := c.DeleteIdentityResolver(resolverConfigId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
