---
subcategory: ""
page_title: "Manage Users/Clients in a CSV - Unifi Provider"
description: |-
An example of using a CSV to manage all of your users of your network.
---


You could create/manage a `unifi_user` for every row/MAC address in the CSV with the following config:

{{ tffile "examples/data-sources/subscription_congih.tf" }}