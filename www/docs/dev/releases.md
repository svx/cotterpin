---
title: Releases
id: releases
description: How to make releases
sidebar_position: 2
---

How we do releases with GoReleaser and GitHub Actions.

---

Our Git repository is hosted in GitHub so we will use GitHub Actions for our CI (Continuous Integration) pipeline.

Create our workflow

```shell
mkdir .github/workflows
cd .github/workflows
```

Inside it, create goreleaser.yml file with this content:

```yml
name: goreleaser

on:
  push:
    tags:
      - '*'

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v2
        with:
          fetch-depth: 0
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.17
      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v2
        with:
          distribution: goreleaser
          version: latest
          args: release --rm-dist
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

This workflow contains one job that we'll checkout the repository, package our app with GoReleaser and generate a GitHub release.

:::note
To release to GitHub, GoReleaser need a valid GitHub token with the repo scope.
Fortunately, GitHub automatically [creates a GITHUB_TOKEN](https://docs.github.com/en/actions/reference/authentication-in-a-workflow#about-the-github_token-secret) secret to use in a workflow.
:::

After pushed your modification in the Git repository, now we can create a Git tag and push it:

```shell
git tag -a v1.0.0 -m "First release"

git push --tags
To https://github.com/svx/cotterpin.git
 * [new tag]         v1.0.0 -> v1.0.0
```

Each time we will update the app and create a Git tag and push it, automatically a new (GitHub) GH release will be created with cross-platform binaries :-).
