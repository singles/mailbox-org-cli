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
    "expires": "2022-02-28"
  },
  {
    "email": "aks92jasl943@temp.mailbox.org",
    "memo": "", # there's no memo set
    "expires": "2022-03-31"
  }
]
```

All output is JSON, so you will probably need something like [`jq`](https://github.com/stedolan/jq) to extract specific data. Using example output above this command will copy first item's email into clipboard (`pbcopy` on macOS):

```shell
$ mailbox-org-cli ... list | jq --raw '.[0].email' | pbcopy
```

## Building

You need [Go](https://go.dev/) in version at least `1.17`. In the project's root directory run `go build .`.
