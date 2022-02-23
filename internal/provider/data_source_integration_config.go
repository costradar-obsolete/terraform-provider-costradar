package provider

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var integrationConfigSchema = map[string]*schema.Schema{
	"integration_role_arn": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"integration_sqs_url": {
		Type: schema.TypeString,
		Computed: true,
	},
	"integration_sqs_arn": {
		Type: schema.TypeString,
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
	d.Set("integration_role_arn", integrationConfig.IntegrationRoleArn)
	d.Set("integration_sqs_url", integrationConfig.IntegrationSqsUrl)
	d.Set("integration_sqs_arn", integrationConfig.IntegrationSqsArn)

	id := integrationConfig.IntegrationSqsArn + integrationConfig.IntegrationSqsUrl + integrationConfig.IntegrationRoleArn
	hasher := sha1.New()
	hasher.Write([]byte(id))
	shaId := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	d.SetId(shaId)
	return diags
}
