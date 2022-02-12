package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var resourceAccountSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"account_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"alias": {
		Type:     schema.TypeString,
		Required: true,
	},
	"owners": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
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
	"tags": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}

func resourceAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAccountCreate,
		ReadContext:   resourceAccountRead,
		UpdateContext: resourceAccountUpdate,
		DeleteContext: resourceAccountDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceAccountSchema,
	}
}

func accountFromResourceData(d *schema.ResourceData) Account {
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
	account := Account{
		ID:           d.Get("id").(string),
		AccountId:    d.Get("account_id").(string),
		Alias:        d.Get("alias").(string),
		AccessConfig: accessConfig,
		Owners:       d.Get("owners").([]interface{}),
		Tags:         d.Get("tags").(map[string]interface{}),
	}

	return account
}

func accountToResourceData(d *schema.ResourceData, a Account) {
	var accessConfigList []map[string]string
	accessConfig := make(map[string]string)
	accessConfig["reader_mode"] = a.AccessConfig.ReaderMode
	accessConfig["assume_role_arn"] = a.AccessConfig.AssumeRoleArn
	accessConfig["assume_role_external_id"] = a.AccessConfig.AssumeRoleExternalId
	accessConfig["assume_role_session_name"] = a.AccessConfig.AssumeRoleSessionName
	accessConfigList = append(accessConfigList, accessConfig)
	d.Set("account_id", a.AccountId)
	d.Set("alias", a.Alias)
	d.Set("access_config", accessConfigList)
	d.Set("owners", a.Owners)
	d.Set("tags", a.Tags)
}

func resourceAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	id := d.Id()
	account, err := c.GetAccount(id)
	if err != nil {
		return diag.FromErr(err)
	}
	accountToResourceData(d, account.Payload)
	return diags
}

func resourceAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	account := accountFromResourceData(d)

	a, err := c.CreateAccount(account)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(a.Payload.ID)
	resourceAccountRead(ctx, d, m)
	return diags
}

func resourceAccountUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	account := accountFromResourceData(d)
	c := m.(Client)
	_, err := c.UpdateAccount(account)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceAccountRead(ctx, d, m)
}

func resourceAccountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteAccount(id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
