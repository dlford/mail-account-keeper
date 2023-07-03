# Mail Account Keeper

A daemon written in Go for sending periodic emails from an account so it doesn't get deleted.

## Getting Started

Create a JSON array of the accounts you'd like to preserve, in the following syntax:

```json
[
  {
    "title": "A name for this account",
    "host": "e.g. smtp.your-mail-host.com",
    "port": 587, // Choose your mail port
    "email": "you@your-mail-host.com",
    "password": "your secret password (better to create an app password where possible)",
    "mail_to": "email address to send the messages to",
    "schedule": "0 0 1 * *" // CRON syntax
  }
  // Add as many accounts here as needed
]
```

You can pass this as an argument to the executable (`--accounts "[...]"`), or as a Docker environemnt variable (`-e MAIL_ACCOUNT_KEEPER_ACCOUNTS="[...]"`).
