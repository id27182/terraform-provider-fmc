---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_host_objects Data Source - fmc-terraform"
subcategory: ""
description: |-
  Data source for Host Objects in FMC
  An example is shown below:
  hcl
  data "fmc_host_objects" "existing_host" {
      name = "CUCM-Pub"
  }
  
  Any one of the id, name or value can be specified. The first filter in the order of id, name and value will be used, and the rest will be ignored if multiple are specified.
---

# fmc_host_objects (Data Source)

Data source for Host Objects in FMC

An example is shown below: 
```hcl
data "fmc_host_objects" "existing_host" {
	name = "CUCM-Pub"
}
```
Any one of the id, name or value can be specified. The first filter in the order of id, name and value will be used, and the rest will be ignored if multiple are specified.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource
- **name** (String) The name of this resource
- **value** (String) The value of this resource

### Read-Only

- **type** (String) The type of this resource


