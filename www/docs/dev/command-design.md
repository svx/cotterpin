---
title: Command design
id: command-design
description: Command design
sidebar_position: 5
---

Here a description, or we could add this part to the style-guide?

---

For nontrivial commands there’s a hierarchy of these nouns and verbs that users put together to solve a problem.
Either use a verb-object construction like `cotterpin add readme` , where `add` is a top-level verb and the targets vary.

Or use subject-verb construction like `gcloud container clusters list` (“clusters” is a subcommand of “container” and “list” is the verb).

```ini
Available Commands:
  add         Add a part of a component
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initialize project components
  version     Show version
```
