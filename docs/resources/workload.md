---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "costradar_workload Resource - terraform-provider-costradar"
subcategory: "Workloads"
description: |-
  
---

# costradar_workload (Resource)



## Example Usage

```terraform
resource "costradar_workload" "this" {
  name = "Delivery Service"
  description = "Core delivery workload"
  owners = [
    "admin@email.com"
  ]
  tags = {
    "unit": "Delivery"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String)

### Optional

- **description** (String)
- **owners** (List of String)
- **tags** (Map of String)

### Read-Only

- **id** (String) The ID of this resource.


