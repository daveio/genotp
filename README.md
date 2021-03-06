# `gotp`

Generate OATH-TOTP one-time passwords from the command line.

<table>
  <thead>
    <tr>
      <th>
        <code>master</code>
      </th>
      <th>
        <code>gh-pages</code>
      </th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>
        <a href="https://travis-ci.com/daveio/gotp/branches" rel="nofollow">
          <img src="https://travis-ci.com/daveio/gotp.svg?branch=master" alt="master branch build status">
        </a>
      </td>
      <td>
        <a href="https://travis-ci.com/daveio/gotp/branches" rel="nofollow">
          <img src="https://travis-ci.com/daveio/gotp.svg?branch=gh-pages" alt="gh-pages branch build status">
        </a>
      </td>
    </tr>
  </tbody>
</table>

## Installation

### The easy way, with Homebrew

With Homebrew installed, first add my personal tap.

`brew tap daveio/daveio`

You only need to do this once, after which all of my projects will then be available for installation, and Homebrew will find the latest versions of my software on an ongoing basis.

Once the tap is added, simply do

`brew install gotp`

and you're ready to roll.

### The easy way, with Go

If you have a working Go installation, all you need to do is

`go get github.com/daveio/gotp`

after which you'll have a shiny new `gotp` binary in your `$GOPATH/bin`.

### The easy way, with `zsh`

I also develop [`zsh-gotp`][link-zsh-gotp], a `zsh` plugin which handles automatic installation and setup of aliases and completion. If you use `zsh` it's strongly recommended and might save you a lot of effort.

### The manual way

You can also download a standalone binary from this repository's [Releases page][link-gotp-releases].

Currently, binaries are available for macOS (amd64 only), Linux (i386 and amd64), and Windows (i386 and amd64). Put the `gotp` (or `gotp.exe` for Windows) binary somewhere in your `$PATH` and you're done.

If you want additional architectures added to the build scripts, [open a feature request Issue][link-open-feature-request] and let me know. Accompanying the Issue with a pull request with relevant modifications to the build script is the best way to get it live quickly.

## Integration

`gotp` will work just fine on its own, but there are a few ways to reduce friction even further.

### zsh plugin

Try [`zsh-gotp`][link-zsh-gotp].

### Shell integration

*The following contains clipboard functionality specific to macOS, but is easily adapted to other systems.*

If you're using [`zsh-gotp`][link-zsh-gotp] then this will be automatically set up for you, but if you want to do it manually, add the following function to your shell's rc file:

```sh
otp() {
  out=$(gotp generate ${1})
  pwd=$(echo "${out}" | cut -d ":" -f 2 | cut -b 2-)
  echo "${pwd}"
  echo -n "${pwd}" | pbcopy
}
```

You can then do

`otp sitename`

to generate an OTP for the default account for `sitename`, and automatically copy it to the clipboard.

If you want to integrate clipboard functionality on non-macOS systems, find a command which writes `STDIN` to the clipboard and replace `pbcopy` in the function with that command. Alternatively, feel free to comment out the last line entirely, and just copy the output manually.

### Short forms

Each command in `gotp` has a short form. These are listed in `gotp --help`.

|Long form        |Short form|
|-----------------|----------|
|`gotp generate`  |`gotp g`  |
|`gotp store`     |`gotp s`  |
|`gotp delete`    |`gotp d`  |
|`gotp list-sites`|`gotp ls` |
|`gotp list-uids` |`gotp lu` |

## Usage

```text
usage: gotp [<flags>] <command> [<args> ...]

Generate OATH-TOTP one-time passwords from the command line.

Flags:
      --help     Show context-sensitive help (also try --help-long and --help-man).
  -v, --verbose  Show more detail.
      --version  Show application version.

Commands:
  help [<command>...]
    Show help.

  store [<flags>] <site> <key>
    Short form: 's'. Store a new account.

  generate [<flags>] <site>
    Short form: 'g'. Generate OTP(s) for a site.

  delete [<flags>] <site>
    Short form: 'd'. Delete a site or account.

  list-sites
    Short form: 'ls'. List the sites you have added keys for.

  list-uids <site>
    Short form: 'lu'. List the accounts you have added for a site.
```

### Examples

#### Store a new default account for a new site `sitename`

`gotp store sitename KEY123123123123`

#### Generate an OTP for the default account on site `sitename`

`gotp generate sitename`

#### Store a named account for site `sitename`

`gotp store -u accountname sitename KEY123123123123`

#### Generate an OTP for the account `accountname` on the site `sitename`

`gotp generate -u accountname sitename`

#### Delete all accounts for site `sitename`

`gotp delete sitename`

#### Delete a specific account `accountname` on site `sitename`

`gotp delete -u accountname sitename`

#### List the sites you have data for

`gotp list-sites`

#### List the account names for site `sitename`

`gotp list-uids sitename`

## Planned features

### Credential security hardening

Currently, credentials are stored in plain text in a JSON file named `keychain.json`. Also, the default permissions for the file may allow reading by other users on the same system. This situation is suboptimal, to put it lightly.

### QR code decoding

TOTP credentials are usually supplied in the form of a QR code for scanning on a mobile authenticator. The ability to feed `gotp` a screenshot or other image containing a QR code would make the process of getting credentials imported a lot cleaner.

### `totp://` URL parsing

After the QR code is decoded, the actual TOTP credentials are supplied in the form of a URL with the `totp://` scheme. Parsing these URLs natively takes another manual step out of the process.

## Known issues

* Credentials are stored in plain text, and without any specifically strong permissions.
* The internal representation for a site's default credentials (uid `__default`) is exposed to the user.

[link-zsh-gotp]: https://github.com/daveio/zsh-gotp
[link-gotp-releases]: https://github.com/daveio/gotp/releases
[link-open-feature-request]: https://github.com/daveio/gotp/issues/new?assignees=&labels=&template=feature_request.md&title=
