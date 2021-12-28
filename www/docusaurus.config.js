const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').DocusaurusConfig} */
module.exports = {
  title: 'Cotterpin',
  tagline: 'Docs are cool',
  url: 'https://sven.io',
  baseUrl: '/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.png',
  customFields: {
    description:
      'An optimized site generator in React. Docusaurus helps you to move fast and write content. Build documentation websites, blogs, marketing pages, and more.',
  },
  organizationName: 'svx', // Usually your GitHub org/user name.
  projectName: 'cotterpin', // Usually your repo name.
  // plugins: [
  //   require.resolve('@cmfcmf/docusaurus-search-local')
  // ],
  themeConfig: {
    // announcementBar: {
    //   id: 'support_us', // Any value that will identify this message.
    //   content:
    //     'We are looking to revamp our docs, please fill <a target="_blank" rel="noopener noreferrer" href="#">this survey</a>',
    //   backgroundColor: '#fafbfc', // Defaults to `#fff`.
    //   textColor: '#091E42', // Defaults to `#000`.
    //   isCloseable: false, // Defaults to `true`.
    // },
    colorMode: {
      // "light" | "dark"
      defaultMode: 'light',
      // Hides the switch in the navbar
      // Useful if you want to support a single color mode
      disableSwitch: true,
    },
    //hideableSidebar: true,
    prism: {
      theme: require('prism-react-renderer/themes/dracula'),
      additionalLanguages: ['ini', 'graphql', 'git', 'docker', 'makefile'],
    },
    navbar: {
      title: 'Cotterpin',
      logo: {
        alt: 'Logo',
        src: 'img/ocld-logo.png',
      },
      items: [
        {
          //type: 'doc',
          //docId: 'upload',
          to: 'intro',
          position: 'left',
          label: 'Tutorial',
        },
        {
          to: 'blog',
          position: 'left',
          label: 'Blog',
        },
        {
          type: 'doc',
          docId: 'intro',
          label: 'Download',
          position: 'right',
          className: 'btn-light',
        },
        {
          href: 'https://github.com/svx/cotterpin',
          position: 'right',
          className: 'header-github-link',
          'aria-label': 'GitHub repository',
        },
        //{to: '/help', label: 'Support', position: 'left'},

      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Tutorial',
              to: '/intro',
            },
            {
              label: 'Blog',
              to: '/blog',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Stack Overflow',
              href: 'https://stackoverflow.com/questions/tagged/docusaurus',
            },
            {
              label: 'Slack',
              href: 'https://discordapp.com/invite/docusaurus',
            },
            {
              label: 'Twitter',
              href: 'https://twitter.com/docusaurus',
            },
            {
              label: 'Something',
              href: 'https://stackoverflow.com/questions/tagged/docusaurus',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'GitHub',
              href: 'https://github.com/ocular-d',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} ocular-d. Built with Docusaurus.`,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          // editUrl:
          //   'https://github.com/facebook/docusaurus/edit/master/website/',
          showLastUpdateTime: true,
          routeBasePath: '/'
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          path: 'blog',
          routeBasePath: 'blog',
          blogSidebarTitle: 'Documentation updates',
          blogTitle: 'Docs updates',
          editUrl:
            'https://github.com/facebook/docusaurus/edit/master/website/blog/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
};
