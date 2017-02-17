# Mondas

[![Build Status][build-status-image]][build-status-url]
[![Coverage Status][coverage-status-image]][coverage-status-url]
[![Documentation][documentation-image]][documentation-url]
[![Report Card][report-card-image]][report-card-url]

Mondas is a toolkit for managing, packaging, and sharing your internal tools together.

## Prerequisites

- [Go 1.8](https://golang.org/dl/) or higher

## Installation

```
$ go get -u github.com/jamieconnolly/mondas
```

## Usage

```go
package main

import "github.com/jamieconnolly/mondas"

func main() {
    mondas.Run("example", "1.2.3")
}
```

[build-status-image]: https://api.travis-ci.org/jamieconnolly/mondas.svg?branch=master
[build-status-url]: https://travis-ci.org/jamieconnolly/mondas

[coverage-status-image]: https://coveralls.io/repos/github/jamieconnolly/mondas/badge.svg?branch=master
[coverage-status-url]: https://coveralls.io/github/jamieconnolly/mondas?branch=master

[documentation-image]: https://godoc.org/github.com/jamieconnolly/mondas?status.svg
[documentation-url]: https://godoc.org/github.com/jamieconnolly/mondas

[report-card-image]: https://goreportcard.com/badge/jamieconnolly/mondas
[report-card-url]: https://goreportcard.com/report/jamieconnolly/mondas
