package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

var subscriptionMetaSchema = map[string]*schema.Schema{
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
	d.Set("cur_sqs_arn", subscriptionMeta.CurSqsArn)
	d.Set("cur_sqs_url", subscriptionMeta.CurSqsUrl)
	d.Set("cloudtrail_sqs_arn", subscriptionMeta.CloudTrailSqsArn)
	d.Set("cloudtrail_sqs_url", subscriptionMeta.CloudTrailSqsUrl)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
