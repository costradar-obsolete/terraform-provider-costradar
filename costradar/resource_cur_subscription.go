package costradar

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var resourceCurSubscriptionSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"report_name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"bucket_name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"bucket_region": {
		Type:     schema.TypeString,
		Required: true,
	},
	"bucket_path_prefix": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"time_unit": {
		Type:     schema.TypeString,
		Required: true,
	},
	"source_topic_arn": {
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

func resourceCurSubscription() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCurSubscriptionCreate,
		ReadContext:   resourceCurSubscriptionRead,
		UpdateContext: resourceCurSubscriptionUpdate,
		DeleteContext: resourceCurSubscriptionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceCurSubscriptionSchema,
	}
}

func curSubscriptionFromResourceData(d *schema.ResourceData) CostAndUsageReportSubscription {
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
	subscription := CostAndUsageReportSubscription{
		ID:               d.Get("id").(string),
		ReportName:       d.Get("report_name").(string),
		BucketName:       d.Get("bucket_name").(string),
		BucketRegion:     d.Get("bucket_region").(string),
		BucketPathPrefix: d.Get("bucket_path_prefix").(string),
		SourceTopicArn:   d.Get("source_topic_arn").(string),
		TimeUnit:         d.Get("time_unit").(string),
		AccessConfig:     accessConfig,
	}

	return subscription
}

func curSubscriptionToResourceData(d *schema.ResourceData, s CostAndUsageReportSubscription) {
	var accessConfigList []map[string]string
	accessConfig := make(map[string]string)
	accessConfig["reader_mode"] = s.AccessConfig.ReaderMode
	accessConfig["assume_role_arn"] = s.AccessConfig.AssumeRoleArn
	accessConfig["assume_role_external_id"] = s.AccessConfig.AssumeRoleExternalId
	accessConfig["assume_role_session_name"] = s.AccessConfig.AssumeRoleSessionName
	accessConfigList = append(accessConfigList, accessConfig)
	d.Set("report_name", s.ReportName)
	d.Set("bucket_name", s.BucketName)
	d.Set("bucket_region", s.BucketRegion)
	d.Set("bucket_path_prefix", s.BucketPathPrefix)
	d.Set("source_topic_arn", s.SourceTopicArn)
	d.Set("time_unit", s.TimeUnit)
	d.Set("access_config", accessConfigList)
}

func resourceCurSubscriptionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	var subscription = curSubscriptionFromResourceData(d)

	s, err := c.CreateCostAndUsageReportSubscription(subscription)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(s.Payload.ID)
	resourceCurSubscriptionRead(ctx, d, m)

	return diags
}

func resourceCurSubscriptionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	subscriptionId := d.Id()
	subscription, err := c.GetCostAndUsageReportSubscription(subscriptionId)
	if err != nil {
		return diag.FromErr(err)
	}
	curSubscriptionToResourceData(d, subscription.Payload)
	return diags
}

func resourceCurSubscriptionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	subscription := curSubscriptionFromResourceData(d)
	c := m.(Client)
	_, err := c.UpdateCostAndUsageReportSubscription(subscription)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceCurSubscriptionRead(ctx, d, m)
}

func resourceCurSubscriptionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)
	var diags diag.Diagnostics

	subscriptionId := d.Id()

	err := c.DeleteCostAndUsageReportSubscription(subscriptionId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
