---
title: Guide
id: dev-guide
description: Command Messages
sidebar_position: 2
---

Developer overview and guide.

---

This CLI is written in [Cobra](https://github.com/spf13/cobra).
Cobra is a Go library for building CLI applications.

## Developer requirements

- [Golang](https://go.dev/)
- [GoReleaser](https://goreleaser.com/)
- [Task](https://taskfile.dev/#/)

## Getting started


Run `task setup`.

Use `go get` to install the latest version of the library.
This command will install the cobra generator executable along with the library and its dependencies:

```shell
go get -u github.com/spf13/cobra
```

## Commands

Use `cobra add`, the following example creates a new command with the name *foo*.

```shell
cobra add foo
```

### Subcommands

The following example adds the subcommand *bar* to the command *foo*.

```shell
cobra add -p bar fooCmd
```

## Colors

- Green for info
- Yellow for warning
- Red for error

We use the [github.com/fatih/color](https://github.com/fatih/color) library for color support.

```golang {7,10}
import (
	"fmt"
	"os"
	"log"
    "text/template"

	"github.com/fatih/color"
)

color.Green("Processing")
```