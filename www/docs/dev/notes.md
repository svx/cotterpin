---
title: Notes
id: Notes
description: Notes
sidebar_position: 3
---

## Screen recording

We use [Asciinema](https://asciinema.org/)

```shell
asciinema rec cotterpin-0.0.1-dev
```

```shell
docker run --rm -v $PWD:/data asciinema/asciicast2gif cotterpin-0.0.1-dev foobar.gif
==> Loading cotterpin-0.0.1-dev...
==> Spawning PhantomJS renderer...
==> Generating frame screenshots...
==> Combining 21 screenshots into GIF file...
==> Done.
```

Or with solarized dark as theme:

```shell
docker run --rm -v $PWD:/data asciinema/asciicast2gif -s 2 -t solarized-dark cc-01 cc.gif
```

## Go Embed

- https://charly3pins.dev/blog/learn-how-to-use-the-embed-package-in-go-by-building-a-web-page-easily/
- https://medium.com/@leo_hetsch/using-gos-embed-package-to-build-a-small-webpage-6175953fccea
- https://threkk.medium.com/embedding-files-natively-in-go-1-16-2a2f2070617d
- https://github.com/amitsaha/go-embed
- https://echorand.me/posts/go-embed/
- https://blog.jetbrains.com/go/2021/06/09/how-to-use-go-embed-in-go-1-16/
- https://bhupesh-v.github.io/embedding-static-files-in-golang/
- https://www.yellowduck.be/posts/embedding-file-with-go-116/

## Go Copy

- https://golangbyexample.com/copy-file-go/
- https://www.golangprograms.com/golang-create-copy-of-a-file-at-another-location.html

## Go import function

- https://stackoverflow.com/questions/34967358/how-to-import-my-own-function-defined-out-of-main-package-in-go

## Strings

- https://www.digitalocean.com/community/tutorials/how-to-define-and-call-functions-in-go

## Packages/Imports/Functions

- https://www.callicoder.com/golang-packages/
