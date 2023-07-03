# Mail Account Keeper

A daemon written in Go for sending periodic emails from an account so it doesn't get deleted.

## Getting Started

Create a JSON array of the accounts you'd like to preserve, in the following syntax:

```json
[
  {
    "title": "A name for this account",
    "host": "smtp.your-mail-host.com",
    "port": 587,
    "email": "you@your-mail-host.com",
    "username": "username to access your email account, omit if email is username",
    "password": "your secret password (better to create an app password where possible)",
    "mail_to": "email address to send the messages to",
    "schedule": "0 0 1 * *"
  }
]
```

`schedule` should be written in cron syntax.

`username` should be omitted if `email` is your username for logging into the email account.

If you'd like alerts for failed sends, you'll need a JSON object for an alert account:

```json
{
  "host": "smtp.your-mail-host.com",
  "port": 587,
  "email": "you@your-mail-host.com",
  "username": "username to access your email account, omit if email is username",
  "password": "your secret password or app key",
  "mail_to": "you@your-mail-host.com"
}
```

You can pass these as arguments to the executable (`--accounts "[...]" --alerts "{...}"`), or as a Docker environemnt variable (`-e MAIL_ACCOUNT_KEEPER_ACCOUNTS="[...]" -e MAIL_ACCOUNT_KEEPER_ALERTS="{...}"`). Note that "alerts" is optional.

Recommended method is to use Docker.

```sh
docker run -e MAIL_ACCOUNT_KEEPER_ACCOUNTS="[...]" -e MAIL_ACCOUNT_KEEPER_ALERTS="{...}" ghcr.io/dlford/mail-account-keeper:latest
```

## Using with Ansible

You can use YAML syntax with Ansible and pass your config through `to_json`.

Add your config to variables to an Ansible Vault like the following:

```yaml
mail_account_keeper_accounts:
  - title: ...
    host: ...
    port: ...
    email: ...
    password: ...
    mail_to: ...
    schedule: ...
  - title: ...
    host: ...
    port: ...
    email: ...
    password: ...
    mail_to: ...
    schedule: ...
mail_account_keeper_alerts:
  host: ...
  port: ...
  email: ...
  username: ...
  password: ...
  mail_to: ...
```

Then add the task for running mail-account-keeper like this:

```yaml
- name: Start Mail Account Keeper
  become: true
  community.docker.docker_container:
    name: mail-account-keeper
    restart_policy: unless-stopped
    image: ghcr.io/dlford/mail-account-keeper:latest
    env:
      MAIL_ACCOUNT_KEEPER_ACCOUNTS: "{{ mail_account_keeper_accounts | to_json }}"
      MAIL_ACCOUNT_KEEPER_ALERTS: "{{ mail_account_keeper_alerts | to_json }}"
```
