package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("COSTRADAR_TOKEN", nil),
			},
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("COSTRADAR_ENDPOINT", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"costradar_cur_subscription":        resourceCurSubscription(),
			"costradar_cloudtrail_subscription": resourceCloudTrailSubscription(),
			"costradar_identity_resolver":       resourceIdentityResolverConfig(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"costradar_integration_config": dataSourceIntegrationConfig(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	endpoint := d.Get("endpoint").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//diags = append(diags, diag.Diagnostic{
	//	Severity: diag.Warning,
	//	Summary:  "Warning Message Summary",
	//	Detail:   "This is the detailed warning message from providerConfigure",
	//})

	return NewCostRadarClient(endpoint, token), diags
}
