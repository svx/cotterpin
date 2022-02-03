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

## Templates

- https://banzaicloud.com/blog/creating-helm-charts-part-2/
- https://onecompiler.com/posts/3tmdpwmvy/go-language-program-to-replace-spaces-with-hyphen-in-a-string
- https://gist.github.com/Miciah/d35582eebb401f4f9fac1385eb5d7266
- https://groups.google.com/g/golang-nuts/c/6i3u17CkB7E
- https://grafana.com/docs/loki/latest/logql/template_functions/
- https://helm.sh/docs/chart_template_guide/functions_and_pipelines/
- https://pkg.go.dev/text/template#hdr-Pipelines
