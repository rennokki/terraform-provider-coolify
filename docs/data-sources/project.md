---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "coolify_project Data Source - coolify"
subcategory: ""
description: |-
  Get a Coolify project by uuid.
---

# coolify_project (Data Source)

Get a Coolify project by `uuid`.

## Example Usage

```terraform
# Retrieve a specific project
data "coolify_project" "example" {
  uuid = "abc123"
}

output "project_name" {
  value = data.coolify_project.example.name
}

output "project_description" {
  value = data.coolify_project.example.description
}

output "project_environments" {
  value = data.coolify_project.example.environments[*].name
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `uuid` (String) Project UUID

### Read-Only

- `description` (String)
- `environments` (Attributes List) The environments of the project. (see [below for nested schema](#nestedatt--environments))
- `id` (Number) The ID of this resource.
- `name` (String)

<a id="nestedatt--environments"></a>
### Nested Schema for `environments`

Read-Only:

- `created_at` (String)
- `description` (String)
- `id` (Number)
- `name` (String)
- `project_id` (Number)
- `updated_at` (String)
