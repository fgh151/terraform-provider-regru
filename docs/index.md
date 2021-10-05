---
page_title: "Provider: regru dns"
subcategory: "dns"
description: |-
  Terraform provider for interacting with dns records hosted on reg.ru.
---

# Regru DNS Provider

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "regru" {
  password = "passwd"
  username = "username"
}
```

## Schema

### Required

- **username** (String, Optional) Username to authenticate to regru API
- **password** (String, Optional) Password to authenticate to regru API
