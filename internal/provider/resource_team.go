package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var resourceTeamSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"tags": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}

func resourceTeam() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTeamCreate,
		ReadContext:   resourceTeamRead,
		UpdateContext: resourceTeamUpdate,
		DeleteContext: resourceTeamDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceTeamSchema,
	}
}

func TeamFromResourceData(d *schema.ResourceData) Team {
	team := Team{
		ID:          d.Get("id").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Tags:        d.Get("tags").(map[string]interface{}),
	}
	return team
}

func TeamToResourceData(d *schema.ResourceData, t Team) {
	d.Set("name", t.Name)
	if nilIfEmpty(t.Description) != nil {
		d.Set("description", t.Description)
	}
	if nilIfEmpty(t.Tags) != nil {
		d.Set("tags", t.Tags)
	}
}

func resourceTeamRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	id := d.Id()
	Team, err := c.GetTeam(id)
	if err != nil {
		return diag.FromErr(err)
	}
	TeamToResourceData(d, Team.Payload)
	return diags
}

func resourceTeamCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	var Team = TeamFromResourceData(d)

	w, err := c.CreateTeam(Team)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceTeamRead(ctx, d, m)
	d.SetId(w.Payload.ID)
	return diags
}

func resourceTeamUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	Team := TeamFromResourceData(d)
	c := m.(Client)
	_, err := c.UpdateTeam(Team)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceTeamRead(ctx, d, m)
}

func resourceTeamDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)
	var diags diag.Diagnostics

	TeamId := d.Id()

	err := c.DeleteTeam(TeamId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
