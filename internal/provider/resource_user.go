package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

var resourceUserSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"email": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"initials": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"icon_url": {
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

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceUserSchema,
	}
}

func UserFromResourceData(d *schema.ResourceData) User {
	user := User{
		ID:       d.Get("id").(string),
		Email:    d.Get("email").(string),
		Name:     d.Get("name").(string),
		Initials: d.Get("initials").(string),
		IconUrl:  d.Get("icon_url").(string),
		Tags:     d.Get("tags").(map[string]interface{}),
	}
	return user
}

func UserToResourceData(d *schema.ResourceData, u User) {
	d.Set("name", u.Name)
	d.Set("initials", u.Initials)
	d.Set("email", u.Email)
	if nilIfEmpty(u.IconUrl) != nil {
		d.Set("icon_url", u.IconUrl)
	}
	if nilIfEmpty(u.Tags) != nil {
		d.Set("tags", u.Tags)
	}
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var diags diag.Diagnostics

	id := d.Id()
	user, err := c.GetUser(id)

	if user.Payload.ID == "" {
		log.Printf("[WARN] AWS user (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return diag.FromErr(err)
	}
	UserToResourceData(d, user.Payload)
	return diags
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)

	var user = UserFromResourceData(d)

	w, err := c.CreateUser(user.Email, user)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(w.Payload.ID)
	return resourceUserRead(ctx, d, m)
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	user := UserFromResourceData(d)
	c := m.(Client)
	_, err := c.UpdateUser(user)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceUserRead(ctx, d, m)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(Client)
	var diags diag.Diagnostics

	userId := d.Id()

	err := c.DeleteUser(userId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
