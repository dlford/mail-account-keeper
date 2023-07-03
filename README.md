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
    "username": "username to access your email account, use your email address if you don't have one",
    "password": "your secret password (better to create an app password where possible)",
    "mail_to": "email address to send the messages to",
    "schedule": "0 0 1 * *"
  }
]
```

`schedule` should be written in cron syntax.

If you'd like alerts for failed sends, you'll need a JSON object for an alert account:

```json
{
  "host": "smtp.your-mail-host.com",
  "port": 587,
  "email": "you@your-mail-host.com",
  "username": "username to access your email account, use your email address if you don't have one",
  "password": "your secret password or app key",
  "mail_to": "you@your-mail-host.com"
}
```

You can pass these as arguments to the executable (`--accounts "[...]" --alerts "{...}"`), or as a Docker environemnt variable (`-e MAIL_ACCOUNT_KEEPER_ACCOUNTS="[...]" -e MAIL_ACCOUNT_KEEPER_ALERTS="{...}"`). Note that "alerts" is optional.

Recommended method is to use Docker.

```sh
docker run -e MAIL_ACCOUNT_KEEPER_ACCOUNTS="[...]" -e MAIL_ACCOUNT_KEEPER_ALERTS="{...}" ghcr.io/dlford/mail-account-keeper:latest
```
