# mailbox-org-cli

**Unoffical** command line "client" for managing [mailbox.org](https://mailbox.org) [disposable addresses](https://mailbox.org/en/post/more-privacy-with-anonymous-disposable-e-mail-addresses).

TBH it's hard to call it client - it **does not** use API for achieve its goals, as there is no API for that feature.

What it does it pretends to be browser and interacts with management panel in good-old form-based way.

## Installation

You can download prebuilt binary from [Releases](https://github.com/singles/mailbox-org-cli/releases) page, but at the moment macOS binary **isn't** signed & notarized so you will get a warning that application cannot be verified.

If you have Go installed, you can either:
  * install it from source `go install github.com/singles/mailbox-org-cli@latest`
  * clone repository and build it by yourself: `go build .` (requires Go 1.17)


## Usage

```text
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

```text
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

```text
$ mailbox-org-cli ... list | jq --raw '.[0].email' | pbcopy
```

### Possible use cases

* mailbox.org's dispossable addresses have expiry date. But can be extended as many times as required. So if you want to have "permanent" address, just set cron every, lets say, 2 weeks with `mailbox-org-cli renew`.
* as this is CLI tool, you can easily integrate it with some launcher like [Alfred](https://www.alfredapp.com/)
* ...your idea :)
