package costradar

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

var subscriptionMetaSchema = map[string]*schema.Schema{
	"cost_and_usage_report_sqs_arn": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"cost_and_usage_report_sqs_url": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"cloud_trail_sqs_arn": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"cloud_trail_sqs_url": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

func dataSourceSubscriptionMeta() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSubscriptionMetaRead,
		Schema:      subscriptionMetaSchema,
	}
}

func dataSourceSubscriptionMetaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	subscriptionMeta, err := c.GetIntegrationMeta()
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("cost_and_usage_report_sqs_arn", subscriptionMeta.CostAndUsageReportSqsArn)
	d.Set("cost_and_usage_report_sqs_url", subscriptionMeta.CostAndUsageReportSqsUrl)
	d.Set("cloud_trail_sqs_arn", subscriptionMeta.CostAndUsageReportSqsArn)
	d.Set("cloud_trail_sqs_url", subscriptionMeta.CostAndUsageReportSqsUrl)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
