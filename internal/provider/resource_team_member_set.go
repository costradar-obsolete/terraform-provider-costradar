package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

var resourceTeamMemberSetSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"team_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"set_id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"team_member": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"email": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	},
}

func resourceTeamMemberSet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTeamMemberSetCreate,
		ReadContext:   resourceTeamMemberSetRead,
		UpdateContext: resourceTeamMemberSetUpdate,
		DeleteContext: resourceTeamMemberSetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceTeamMemberSetSchema,
	}
}

func TeamMemberSetFromResourceData(d *schema.ResourceData) (string, string, []TeamMemberSetItem) {
	var TeamMemberSet []TeamMemberSetItem
	teamId, setId := d.Get("team_id").(string), ""
	if d.Id() != "" {
		ids := strings.Split(d.Id(), "/")
		teamId, setId = ids[0], ids[1]
	}

	items := d.Get("team_member").(interface{})
	for _, raw := range items.(*schema.Set).List() {
		mapRaw := raw.(map[string]interface{})
		obj := TeamMemberSetItem{
			Email: mapRaw["email"].(string),
		}
		TeamMemberSet = append(TeamMemberSet, obj)
	}
	return teamId, setId, TeamMemberSet
}

func TeamMemberSetToResourceData(d *schema.ResourceData, t []TeamMemberSetItem) {
	var TeamMemberResourceList []map[string]string

	for _, v := range t {
		m := make(map[string]string)
		m["email"] = v.Email
		TeamMemberResourceList = append(TeamMemberResourceList, m)
	}

	d.Set("team_member", TeamMemberResourceList)
	ids := strings.Split(d.Id(), "/")
	teamId := ids[0]
	setId := ids[1]
	d.Set("team_id", teamId)
	d.Set("set_id", setId)
}

func resourceTeamMemberSetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics
	teamId, setId, _ := TeamMemberSetFromResourceData(d)
	TeamMemberSet, err := c.GetTeamMemberSet(teamId, setId)
	if err != nil {
		return diag.FromErr(err)
	}
	TeamMemberSetToResourceData(d, TeamMemberSet.Payload)
	return diags
}

func resourceTeamMemberSetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	teamId, _, TeamMemberSet := TeamMemberSetFromResourceData(d)
	setId := getUniqueId()
	_, err := c.CreateTeamMemberSet(teamId, setId, TeamMemberSet)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(teamId + "/" + setId)

	return resourceTeamMemberSetRead(ctx, d, m)
}

func resourceTeamMemberSetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	teamId, setId, TeamMemberSet := TeamMemberSetFromResourceData(d)
	c := m.(Client)
	_, err := c.UpdateTeamMemberSet(teamId, setId, TeamMemberSet)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceTeamMemberSetRead(ctx, d, m)
}

func resourceTeamMemberSetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)
	var diags diag.Diagnostics

	teamId, setId, _ := TeamMemberSetFromResourceData(d)
	err := c.DeleteTeamMemberSet(teamId, setId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
