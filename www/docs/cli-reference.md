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
➜ ./cotterpin add readme -h
Add a README to a project.

Usage:
  cotterpin add readme [flags]

Flags:
  -f, --force   Backup existent README.md and overwrite it
  -h, --help    help for readme
```

```ini
cotterpin add readme
```

Create a *README.md* from template.

If a file with the name *README.md* already exists the command will fail and `cotterpin`
will display a error message.

```ini
cotterpin add readme --force
```

Use `--force` to overwrite a existing *README.md*.
This command will create a automatically a backup of the *README.md* in the same
directory.

```ini
.
├── README-02-01-2022
├── README.md
└── cotterpin
```

*README-02-01-2022* is a backup of *README.md*

### dockerfile

```ini
cotterpin add dockerfile
```

Add a Dockerfile, you can choose between the following templates

- Debian 11
- Alpine 3.15
- Node 16 on Alpine 3.15

## init

### docs

```ini
cotterpin init docs
```

Create a `/docs` directory.
If the directory already exists, this command will fail and display an error.

### repo

:::info
For the time being only GitHub is supported!
:::

```ini
cotterpin init repo
```

Configure and setup a Git repository with:

- GitHub actions
- Issue and PR templates
- *tbc*

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

Here the docs about version.
