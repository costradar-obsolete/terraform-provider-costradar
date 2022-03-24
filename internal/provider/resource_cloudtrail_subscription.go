package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

var resourceCloudTrailSubscriptionSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"trail_name": {
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

func resourceCloudTrailSubscription() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudTrailSubscriptionCreate,
		ReadContext:   resourceCloudTrailSubscriptionRead,
		UpdateContext: resourceCloudTrailSubscriptionUpdate,
		DeleteContext: resourceCloudTrailSubscriptionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceCloudTrailSubscriptionSchema,
	}
}

func cloudTrailSubscriptionFromResourceData(d *schema.ResourceData) CloudTrailSubscription {
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

	sub := CloudTrailSubscription{
		ID:             d.Get("id").(string),
		TrailName:      d.Get("trail_name").(string),
		BucketName:     d.Get("bucket_name").(string),
		BucketRegion:   d.Get("bucket_region").(string),
		SourceTopicArn: d.Get("source_topic_arn").(string),
		AccessConfig:   accessConfig,
	}

	if v, ok := d.GetOk("bucket_path_prefix"); ok {
		v := v.(string)
		sub.BucketPathPrefix = &v
	}

	return sub
}

func cloudTrailSubscriptionToResourceData(d *schema.ResourceData, s CloudTrailSubscription) {
	var accessConfigList []map[string]string
	accessConfig := make(map[string]string)
	accessConfig["reader_mode"] = s.AccessConfig.ReaderMode
	accessConfig["assume_role_arn"] = s.AccessConfig.AssumeRoleArn
	accessConfig["assume_role_external_id"] = s.AccessConfig.AssumeRoleExternalId
	accessConfig["assume_role_session_name"] = s.AccessConfig.AssumeRoleSessionName
	accessConfigList = append(accessConfigList, accessConfig)
	d.Set("trail_name", s.TrailName)
	d.Set("bucket_name", s.BucketName)
	d.Set("bucket_region", s.BucketRegion)
	d.Set("source_topic_arn", s.SourceTopicArn)
	d.Set("access_config", accessConfigList)
	if s.BucketPathPrefix != nil {
		d.Set("bucket_path_prefix", *s.BucketPathPrefix)
	}
}

func resourceCloudTrailSubscriptionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var subscription = cloudTrailSubscriptionFromResourceData(d)

	s, err := c.CreateCloudTrailSubscription(subscription)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(s.Payload.ID)

	return resourceCloudTrailSubscriptionRead(ctx, d, m)
}

func resourceCloudTrailSubscriptionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	subscriptionId := d.Id()
	subscription, err := c.GetCloudTrailSubscription(subscriptionId)

	if subscription.Payload.ID == "" {
		log.Printf("[WARN] CloudTrail subscription (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return diag.FromErr(err)
	}
	cloudTrailSubscriptionToResourceData(d, subscription.Payload)
	return diags
}

func resourceCloudTrailSubscriptionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	subscription := cloudTrailSubscriptionFromResourceData(d)
	c := m.(Client)
	_, err := c.UpdateCloudTrailSubscription(subscription)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceCloudTrailSubscriptionRead(ctx, d, m)
}

func resourceCloudTrailSubscriptionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)
	var diags diag.Diagnostics

	subscriptionId := d.Id()

	err := c.DeleteCloudTrailSubscription(subscriptionId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
