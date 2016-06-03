# Random String Generator

Generates a cryptographically-secure random character string.

## Installation

Requires Go version 1.6+ for support of the ``vendor`` directory.  Go version 1.5 supports the ``vendor`` directory as well, but the environment variable ``GO15VENDOREXPERIMENT=1`` needs to be set.

```shell
go get github.com/jrnewell/rand-string
```

## Usage

```
SUMMARY:
   rand-string - generates a cryptographically-secure random character string

USAGE:
   rand-string [global options] strLength

VERSION:
   1.0.0

ARGUMENTS:
   strLength    Length of String to Output

GLOBAL OPTIONS:
   --alphaNum, -a   use only alphaNumeric characters
   --filter, -f 'no-filter' valid character filter (regex format)
   --pass, -p     use only password-friendly characters
   --version, -v    print the version
   --help, -h     show help
```
