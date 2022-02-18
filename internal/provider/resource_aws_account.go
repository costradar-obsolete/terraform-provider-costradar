package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var resourceAwsAccountSchema = map[string]*schema.Schema{
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

func resourceAwsAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAwsAccountCreate,
		ReadContext:   resourceAwsAccountRead,
		UpdateContext: resourceAwsAccountUpdate,
		DeleteContext: resourceAwsAccountDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceAwsAccountSchema,
	}
}

func accountAwsFromResourceData(d *schema.ResourceData) Account {
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

func awsAccountToResourceData(d *schema.ResourceData, a Account) {
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
	if nilIfEmpty(a.Owners) != nil {
		d.Set("owners", a.Owners)
	}
	if nilIfEmpty(a.Tags) != nil {
		d.Set("tags", a.Tags)
	}
	//d.Set("owners", a.Owners)
	//d.Set("tags", a.Tags)
}

func resourceAwsAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	id := d.Id()
	account, err := c.GetAwsAccount(id)
	if err != nil {
		return diag.FromErr(err)
	}
	awsAccountToResourceData(d, account.Payload)
	return diags
}

func resourceAwsAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	account := accountAwsFromResourceData(d)

	a, err := c.CreateAwsAccount(account)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceAwsAccountRead(ctx, d, m)
	d.SetId(a.Payload.ID)
	return diags
}

func resourceAwsAccountUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	account := accountAwsFromResourceData(d)
	c := m.(Client)
	_, err := c.UpdateAwsAccount(account)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceAwsAccountRead(ctx, d, m)
}

func resourceAwsAccountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteAwsAccount(id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
