[![Build Status](https://travis-ci.org/pcarrier/gauth.png?branch=master)](https://travis-ci.org/pcarrier/gauth)

gauth: replace Google Authenticator
===================================

---

This fork replaces OpenSSL password encryption with GnuPG encryption.
It only works with encrypted files as a way of helping to prevent storing
plain text files with this sensible information.


This README is almost the same as the original one with just enough
modifications for making it match with how the tool works now.

---


Installation
------------

With a Go environment already set up, it should be as easy as `go get github.com/fervic/gauth`.

*Eg,* with `GOPATH=$HOME/go`, it will create a binary `$HOME/go/bin/gauth`.

Usage
-----

- In web interfaces, pretend you can't read QR codes, get a secret like `hret 3ij7 kaj4 2jzg` instead.
- Store one secret per line in `gauth.csv`, in the format `name:secret`. For example:

        AWS:   ABCDEFGHIJKLMNOPQRSTUVWXYZ234567ABCDEFGHIJKLMNOPQRSTUVWXYZ234567
        Airbnb:abcd efgh ijkl mnop
        Google:a2b3c4d5e6f7g8h9
        Github:234567qrstuvwxyz

- Encrypt the file:

        $ gpg --encrypt --recipient <you> gauth.csv

- Move the file to your home's config folder:

        $ mv gauth.csv.gpg ~/.config

- Restrict access to your user:

        $ chmod 600 ~/.config/gauth.csv.gpg

- Run `gauth`. The progress bar indicates how far the next change is.

        $ gauth
                   prev   curr   next
        AWS        315306 135387 483601
        Airbnb     563728 339206 904549
        Google     453564 477615 356846
        Github     911264 548790 784099
        [=======                      ]

- `gauth` is convenient to use in `watch`.

        $ watch -n1 gauth

- Remember to keep your system clock synchronized and to **lock your computer when brewing your tea!**

Encryption
----------

`gauth` only works with [GnuPG](https://gnupg.org/) encrypted files.

It calls the `gpg` command line utility, making it a requirement. It is
suggested to configure a `gpg-agent` so that pass phrases are cached.

The program should launch the pass phrase prompt that is configured for the
agent.

Compatibility
-------------

Tested with:

- Airbnb
- Apple
- AWS
- DreamHost
- Dropbox
- Evernote
- Facebook
- Gandi
- Github
- Google
- LastPass
- Linode
- Microsoft
- Okta (reported by Bryan Baldwin)
- WP.com

Please report further results to pierre@gcarrier.fr.

Rooted Android?
---------------

If your Android phone is rooted, it's easy to "back up" your secrets from an `adb shell` into `gauth`.

    # sqlite3 /data/data/com.google.android.apps.authenticator2/databases/database \
              'select email,secret from accounts'

Really, does this make sense?
-----------------------------

At least to me, it does. My laptop features encrypted storage, a stronger authentication mechanism,
and I take good care of its physical integrity.

My phone also runs arbitrary apps, is constantly connected to the Internet, gets forgotten on tables.

Thanks to the convenience of a command line utility, my usage of 2-factor authentication went from
3 to 10 services over a few days.

Clearly a win for security.

:+1: from me (fervic)
