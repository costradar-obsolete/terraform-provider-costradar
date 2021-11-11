package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

var integrationConfigSchema = map[string]*schema.Schema{
	"integration_role_arn": {
		Type: schema.TypeString,
		Computed: true,
	},
	"cur_sqs_arn": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"cur_sqs_url": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"cloudtrail_sqs_arn": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"cloudtrail_sqs_url": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

func dataSourceIntegrationConfig() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIntegrationConfigRead,
		Schema:      integrationConfigSchema,
	}
}

func dataSourceIntegrationConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	integrationConfig, err := c.GetIntegrationConfig()
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("cur_sqs_arn", integrationConfig.CurSqsArn)
	d.Set("cur_sqs_url", integrationConfig.CurSqsUrl)
	d.Set("cloudtrail_sqs_arn", integrationConfig.CloudTrailSqsArn)
	d.Set("cloudtrail_sqs_url", integrationConfig.CloudTrailSqsUrl)
	d.Set("integration_role_arn", integrationConfig.IntegrationRoleArn)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
