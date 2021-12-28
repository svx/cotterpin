---
title: CLI Reference
id: cli-reference
description: CLI Reference
sidebar_position: 3
---

This reference documents CLI commands and flags.

---

```ini
Usage:
  cotterpin [command]

Available Commands:
  add         Add a part of a component
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initialize project components
  version     Show version

Flags:
      --config string   config file (default is $HOME/.cotterpin.yaml)
  -d, --debug           verbose logging
  -h, --help            help for cotterpin

Use "cotterpin [command] --help" for more information about a command.
```

## add

```ini
cotterpin add -h
Add a part of a component. For example:
- Dockerfile
- README

Use this if you want to extent your component.

Usage:
  cotterpin add [flags]
  cotterpin add [command]

Available Commands:
  readme      Add a README

Flags:
  -h, --help   help for add

Global Flags:
      --config string   config file (default is $HOME/.cotterpin.yaml)
  -d, --debug           verbose logging

Use "cotterpin add [command] --help" for more information about a command.
```

### readme

```ini
cotterpin add readme
```

Here the docs about this command.

### dockerfile

:::info
Coming soon, not implemented, yet!
:::

```ini
cotterpin add dockerfile
```

## version

```ini
cotterpin version
Version and build Info

Version: dev
Build Date: unknown
Git Commit: none
OS: linux
Arch: amd64
Go Version: go1.17

Please provide above output for support
```

Here the docs about version