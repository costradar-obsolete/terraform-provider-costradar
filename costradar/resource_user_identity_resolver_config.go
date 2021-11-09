package costradar

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

import (
	"context"
)

var resourceUserIdentityResolverConfigSchema = map[string]*schema.Schema{
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

func resourceUserIdentityResolverConfig() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserIdentityResolverConfigCreate,
		ReadContext:   resourceUserIdentityResolverConfigRead,
		UpdateContext: resourceUserIdentityResolverConfigUpdate,
		DeleteContext: resourceUserIdentityResolverConfigDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceUserIdentityResolverConfigSchema,
	}
}

func userIdentityResolverConfigFromResourceData(d *schema.ResourceData) UserIdentityResolverConfig {
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
	resolverConfig := UserIdentityResolverConfig{
		ID:           d.Get("id").(string),
		LambdaArn:    d.Get("lambda_arn").(string),
		AccessConfig: accessConfig,
	}

	return resolverConfig
}

func userIdentityResolverToResourceData(d *schema.ResourceData, rc UserIdentityResolverConfig) {
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

func resourceUserIdentityResolverConfigCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	var resolverConfig = userIdentityResolverConfigFromResourceData(d)

	s, err := c.CreateUserIdentityResolverConfig(resolverConfig)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(s.Payload.ID)
	resourceUserIdentityResolverConfigRead(ctx, d, m)

	return diags
}

func resourceUserIdentityResolverConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	resolverConfigId := d.Id()
	resolverConfig, err := c.GetUserIdentityResolverConfig(resolverConfigId)
	if err != nil {
		return diag.FromErr(err)
	}
	userIdentityResolverToResourceData(d, resolverConfig.Payload)
	return diags
}

func resourceUserIdentityResolverConfigUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	resolverConfig := userIdentityResolverConfigFromResourceData(d)
	c := m.(Client)
	_, err := c.UpdateUserIdentityResolverConfig(resolverConfig)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceUserIdentityResolverConfigRead(ctx, d, m)
}

func resourceUserIdentityResolverConfigDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)
	var diags diag.Diagnostics

	resolverConfigId := d.Id()

	err := c.DeleteUserIdentityResolverConfig(resolverConfigId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
