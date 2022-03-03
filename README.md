# mailbox-org-cli

**Unoffical** command line "client" for managing [mailbox.org](https://mailbox.org) [disposable addresses](https://mailbox.org/en/post/more-privacy-with-anonymous-disposable-e-mail-addresses).

TBH it's hard to call it client - it **does not** use API for achieve its goals, as there is no API for that feature.

What it does it pretends to be browser and interacts with management panel in good-old form-based way.

## Usage

```shell
$ mailbox-org-cli --help
Commad line "client" for mailbox.org dispossable addresses feature
Usage: mailbox-org-cli --username USERNAME <command> [<args>]

Options:
  --username USERNAME    mailbox.org username
  --help, -h             display this help and exit

Commands:
  list                   list dispossable addresses
  renew                  renew dispossable address
  delete                 delete dispossable address
  set-memo               set-memo on existing dispossable address
  create                 create new dispossable address with optional memo
```

`mailbox-org-cli` requires `--username` (like `you@customdomain.com`) and password being passed in via _stdin_.

```shell
$ pass Email/mailbox.org | mailbox-org-cli --username you@example.com

[
  {
    "email": "kajsdlkj230@temp.mailbox.org",
    "memo": "foo bar",
    "expires": "expires 28 Feb, 2022"
  },
  {
    "email": "aks92jasl943@temp.mailbox.org",
    "memo": "<no memo>", # there's no memo set
    "expires": "expires 31 Mar, 2022"
  }
]
```

Note that `mailbox-org-cli` doesn't care how your password will happen to be in _stdin_, it just has to be there.

You can use password manager like above. (see [`pass`](https://www.passwordstore.org/), but any password manager with CLI will do).

You can store password in some file on encrypted volume: `mailbox-org-cli --username you@example.com < /media/ENCRYPTED_VOLUME/path/to/file/with/password`.

You can store password in ENV variable and then: `echo $MAILBOX_ORG_PASS | mailbox-org-cli --username you@example.com`.

Or you can just echo it directly if you like living on the edge: `echo 'mypassword' | mailbox-org-cli --username you@example.com`.

## Building

You need [Go](https://go.dev/) in version at least `1.17`. In the project's root directory run `go build`.
