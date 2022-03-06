# mailbox-org-cli

**Unoffical** command line "client" for managing [mailbox.org](https://mailbox.org) [disposable addresses](https://mailbox.org/en/post/more-privacy-with-anonymous-disposable-e-mail-addresses).

TBH it's hard to call it client - it **does not** use API for achieve its goals, as there is no API for that feature.

What it does it pretends to be browser and interacts with management panel in good-old form-based way.

## Usage

```shell
$ mailbox-org-cli --help
Commad line "client" for mailbox.org dispossable addresses feature
Usage: mailbox-org-cli --username USERNAME [--password PASSWORD] [--password-on-stdin] <command> [<args>]

Options:
  --username USERNAME    mailbox.org username [env: MAILBOX_ORG_USERNAME]
  --password PASSWORD    mailbox.org password [env: MAILBOX_ORG_PASSWORD]
  --password-on-stdin    read password from stdin
  --help, -h             display this help and exit

Commands:
  list                   list dispossable addresses
  renew                  renew dispossable address
  delete                 delete dispossable address
  set-memo               set-memo on existing dispossable address
  create                 create new dispossable address with optional memo
```

Here is an example how you can use this command with password manager:

```shell
$ pass Email/mailbox.org | mailbox-org-cli --username you@example.com --password-on-stdin list

[
  {
    "email": "kajsdlkj230@temp.mailbox.org",
    "memo": "foo bar",
    "expires": "expires 28 Feb, 2022"
  },
  {
    "email": "aks92jasl943@temp.mailbox.org",
    "memo": "", # there's no memo set
    "expires": "expires 31 Mar, 2022"
  }
]
```

## Building

You need [Go](https://go.dev/) in version at least `1.17`. In the project's root directory run `go build .`.
