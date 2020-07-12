# go-chatwords

[![go.dev](https://img.shields.io/badge/go.dev-pkg-007d9c.svg?style=flat)](https://pkg.go.dev/github.com/suzuki-shunsuke/go-chatwords)
[![Build Status](https://github.com/suzuki-shunsuke/go-chatwords/workflows/CI/badge.svg)](https://github.com/suzuki-shunsuke/go-chatwords/actions)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/go-chatwords/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/go-chatwords)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-chatwords)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-chatwords)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-chatwords.svg)](https://github.com/suzuki-shunsuke/go-chatwords)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-chatwords/master/LICENSE)

Parse line as chat words.

This package is inspired to [github.com/mattn/go-shellwords](https://github.com/mattn/go-shellwords).

## Motivation

To develop a chat bot, it is useful to split a string into substrings.
[There are many awesome packages for cli tool](https://github.com/avelino/awesome-go#command-line), but it is out of scope for many of them to split a string into substrings.
And some existing packages are made for shell, so they are inapropriate for the chat bot in terms of treat of special characters.

## Usage

```golang
// ["<UXXXXX>", "echo", "hello world"], ""
args, text := chatwords.Split("<UXXXXX> echo 'hello world'", -1)
// ["<UXXXXX>", "echo"], " 'hello world'"
args, text := chatwords.Split("<UXXXXX> echo 'hello world'", 2)
```

[The test code](split_test.go) is also useful to understand this package's behavior.

## Rules to split

### special characters

* quote: `'`, `"`
* space: ` `, `\t`
* escape: `\\`

This package is not for shell, so the number of special characters is very small.

## License

[MIT](LICENSE)
