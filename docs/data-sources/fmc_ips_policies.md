---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_ips_policies Data Source - fmc-terraform"
subcategory: ""
description: |-
  Data source for IPS Policy in FMC
  An example is shown below:
  hcl
  data "fmc_ips_policies" "ips_policy" {
      name = "Connectivity Over Security"
  }
---

# fmc_ips_policies (Data Source)

Data source for IPS Policy in FMC

An example is shown below: 
```hcl
data "fmc_ips_policies" "ips_policy" {
	name = "Connectivity Over Security"
}
```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) Name of the IPS policy

### Read-Only

- **id** (String) The ID of this resource
- **type** (String) The type of this resource


