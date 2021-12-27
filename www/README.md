# Website

This website is built using [Docusaurus 2](https://docusaurus.io/), a modern static website generator.

URL: https://elated-ritchie-5bd94d.netlify.app/

### Plugins

- https://github.com/cmfcmf/docusaurus-search-local

### Installation

```
$ yarn
```

### Local Development

```
$ yarn start
```

This command starts a local development server and opens up a browser window. Most changes are reflected live without having to restart the server.

### Build

```
$ yarn build
```

This command generates static content into the `build` directory and can be served using any static contents hosting service.

### Deployment

```
$ GIT_USER=<Your GitHub username> USE_SSH=true yarn deploy
```

If you are using GitHub pages for hosting, this command is a convenient way to build the website and push to the `gh-pages` branch.

## Notes

- https://github.com/facebook/docusaurus/blob/c8cf48a355b01356c0be541a4c16a1907ad0bc3a/website/docusaurus.config.js#L139-L147
- https://docusaurus.io/docs/api/plugins/@docusaurus/plugin-ideal-image
- https://docusaurus.io/docs/api/plugins
