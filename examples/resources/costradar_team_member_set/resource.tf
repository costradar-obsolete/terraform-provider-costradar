resource "costradar_team" "this" {
  name = "Costradar Team"
}


resource "costradar_team_member_set" "this" {
  team_id = costradar_team.this.id

  team_member {
    email = "user1@gmail.com"
  }

  team_member {
    email = "user2@gmail.com"
  }
}