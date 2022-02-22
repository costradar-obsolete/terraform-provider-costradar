package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

var resourceUserIdentitySetSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"user_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"set_id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"user_identity": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"service_vendor": {
					Type:     schema.TypeString,
					Required: true,
				},
				"identity": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	},
}

func resourceUserIdentitySet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserIdentitySetCreate,
		ReadContext:   resourceUserIdentitySetRead,
		UpdateContext: resourceUserIdentitySetUpdate,
		DeleteContext: resourceUserIdentitySetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceUserIdentitySetSchema,
	}
}

func userIdentitySetFromResourceData(d *schema.ResourceData) (string, string, []UserIdentitySetItem) {
	var userIdentitySet []UserIdentitySetItem
	userId, setId := d.Get("user_id").(string), ""
	if d.Id() != "" {
		ids := strings.Split(d.Id(), "/")
		userId, setId = ids[0], ids[1]
	}

	items := d.Get("user_identity").(interface{})
	for _, raw := range items.(*schema.Set).List() {
		mapRaw := raw.(map[string]interface{})
		obj := UserIdentitySetItem{
			Identity:      mapRaw["identity"].(string),
			ServiceVendor: mapRaw["service_vendor"].(string),
		}
		userIdentitySet = append(userIdentitySet, obj)
	}
	return userId, setId, userIdentitySet
}

func userIdentitySetToResourceData(d *schema.ResourceData, i []UserIdentitySetItem) {
	var UserIdentityList []map[string]string

	for _, v := range i {
		m := make(map[string]string)
		m["identity"] = v.Identity
		m["service_vendor"] = v.ServiceVendor
		UserIdentityList = append(UserIdentityList, m)
	}

	d.Set("user_identity", UserIdentityList)
	ids := strings.Split(d.Id(), "/")
	userId := ids[0]
	setId := ids[1]
	d.Set("user_id", userId)
	d.Set("set_id", setId)
}

func resourceUserIdentitySetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics
	userId, setId, _ := userIdentitySetFromResourceData(d)
	userIdentitySet, err := c.GetUserIdentitySet(userId, setId)
	if err != nil {
		return diag.FromErr(err)
	}
	userIdentitySetToResourceData(d, userIdentitySet.Payload)
	return diags
}

func resourceUserIdentitySetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	userId, _, userIdentitySet := userIdentitySetFromResourceData(d)
	setId := getUniqueId()
	_, err := c.CreateUserIdentitySet(userId, setId, userIdentitySet)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(userId + "/" + setId)

	return resourceUserIdentitySetRead(ctx, d, m)
}

func resourceUserIdentitySetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	userId, setId, userIdentitySet := userIdentitySetFromResourceData(d)
	c := m.(Client)
	_, err := c.UpdateUserIdentitySet(userId, setId, userIdentitySet)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceUserIdentitySetRead(ctx, d, m)
}

func resourceUserIdentitySetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)
	var diags diag.Diagnostics

	userId, setId, _ := userIdentitySetFromResourceData(d)
	err := c.DeleteUserIdentitySet(userId, setId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
