---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "coolify_teams Data Source - coolify"
subcategory: ""
description: |-
  Get a list of Coolify teams.
---

# coolify_teams (Data Source)

Get a list of Coolify teams.

## Example Usage

```terraform
# Retrieve all private keys
data "coolify_teams" "all" {}

# Retrieve private keys with a specific description and team_id
data "coolify_teams" "filtered" {
  filter {
    name   = "discord_enabled"
    values = ["true"]
  }
  # (AND)
  filter {
    name   = "id"
    values = ["0", "1"] # (OR)
  }
}

output "all" {
  value = nonsensitive(data.coolify_teams.all.teams[*].name)
}

output "filtered" {
  value = nonsensitive(data.coolify_teams.filtered.teams[*].name)
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (Block List) Filter results by values (see [below for nested schema](#nestedblock--filter))
- `with_members` (Boolean) Whether to fetch team members. This requires an additional API call per team.

### Read-Only

- `teams` (Attributes Set) (see [below for nested schema](#nestedatt--teams))

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Required:

- `name` (String) Name of the field to filter on. Valid names are `name`, `description`, `id`, `discord_enabled`, `resend_enabled`, `smtp_enabled`, `telegram_enabled`
- `values` (List of String) List of values to match against - if any value matches, the filter is satisfied (**OR** operation). Non-string values will be converted to strings if possible, ie `true` -> `"true"`


<a id="nestedatt--teams"></a>
### Nested Schema for `teams`

Read-Only:

- `created_at` (String) The date and time the team was created.
- `custom_server_limit` (String) The custom server limit.
- `description` (String) The description of the team.
- `discord_enabled` (Boolean) Whether Discord is enabled or not.
- `discord_notifications_database_backups` (Boolean) Whether to send database backup notifications via Discord.
- `discord_notifications_deployments` (Boolean) Whether to send deployment notifications via Discord.
- `discord_notifications_scheduled_tasks` (Boolean) Whether to send scheduled task notifications via Discord.
- `discord_notifications_status_changes` (Boolean) Whether to send status change notifications via Discord.
- `discord_notifications_test` (Boolean) Whether to send test notifications via Discord.
- `discord_webhook_url` (String, Sensitive) The Discord webhook URL.
- `id` (Number) The unique identifier of the team.
- `members` (Attributes List) The members of the team. (see [below for nested schema](#nestedatt--teams--members))
- `name` (String) The name of the team.
- `personal_team` (Boolean) Whether the team is personal or not.
- `resend_api_key` (String, Sensitive) The resending API key.
- `resend_enabled` (Boolean) Whether to enable resending or not.
- `show_boarding` (Boolean) Whether to show the boarding screen or not.
- `smtp_enabled` (Boolean) Whether SMTP is enabled or not.
- `smtp_encryption` (String) The SMTP encryption.
- `smtp_from_address` (String) The email address to send emails from.
- `smtp_from_name` (String) The name to send emails from.
- `smtp_host` (String) The SMTP host.
- `smtp_notifications_database_backups` (Boolean) Whether to send database backup notifications via SMTP.
- `smtp_notifications_deployments` (Boolean) Whether to send deployment notifications via SMTP.
- `smtp_notifications_scheduled_tasks` (Boolean) Whether to send scheduled task notifications via SMTP.
- `smtp_notifications_status_changes` (Boolean) Whether to send status change notifications via SMTP.
- `smtp_notifications_test` (Boolean) Whether to send test notifications via SMTP.
- `smtp_password` (String, Sensitive) The SMTP password.
- `smtp_port` (String) The SMTP port.
- `smtp_recipients` (String) The email addresses to send emails to.
- `smtp_timeout` (String) The SMTP timeout.
- `smtp_username` (String) The SMTP username.
- `telegram_chat_id` (String) The Telegram chat ID.
- `telegram_enabled` (Boolean) Whether Telegram is enabled or not.
- `telegram_notifications_database_backups` (Boolean) Whether to send database backup notifications via Telegram.
- `telegram_notifications_database_backups_message_thread_id` (String) The Telegram database backup message thread ID.
- `telegram_notifications_deployments` (Boolean) Whether to send deployment notifications via Telegram.
- `telegram_notifications_deployments_message_thread_id` (String) The Telegram deployment message thread ID.
- `telegram_notifications_scheduled_tasks` (Boolean) Whether to send scheduled task notifications via Telegram.
- `telegram_notifications_scheduled_tasks_thread_id` (String) The Telegram scheduled task message thread ID.
- `telegram_notifications_status_changes` (Boolean) Whether to send status change notifications via Telegram.
- `telegram_notifications_status_changes_message_thread_id` (String) The Telegram status change message thread ID.
- `telegram_notifications_test` (Boolean) Whether to send test notifications via Telegram.
- `telegram_notifications_test_message_thread_id` (String) The Telegram test message thread ID.
- `telegram_token` (String, Sensitive) The Telegram token.
- `updated_at` (String) The date and time the team was last updated.
- `use_instance_email_settings` (Boolean) Whether to use instance email settings or not.

<a id="nestedatt--teams--members"></a>
### Nested Schema for `teams.members`

Read-Only:

- `created_at` (String) The date when the user was created.
- `email` (String) The user email.
- `email_verified_at` (String) The date when the user email was verified.
- `force_password_reset` (Boolean) The flag to force the user to reset the password.
- `id` (Number) The user identifier in the database.
- `marketing_emails` (Boolean) The flag to receive marketing emails.
- `name` (String) The user name.
- `two_factor_confirmed_at` (String) The date when the user two factor was confirmed.
- `updated_at` (String) The date when the user was updated.