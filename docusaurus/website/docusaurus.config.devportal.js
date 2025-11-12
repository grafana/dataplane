// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const { grafanaPrismTheme } = require("./src/theme/prism");

const devPortalHome = "https://grafana-dev.com/developers";

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: "Grafana Data Structure",
  tagline: "A contract of data types as the source of truth",
  url: "https://grafana-dev.com/",
  baseUrl: "developers/dataplane/",
  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",
  favicon: "img/favicon.png",
  organizationName: "grafana",
  projectName: "dataplane",
  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  plugins: [
    [
      "docusaurus-lunr-search",
      {
        disableVersioning: true,
      },
    ],
  ],

  presets: [
    [
      "classic",
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          routeBasePath: "/",
          path: "../../docs/contract",
          sidebarPath: require.resolve("./sidebars.js"),
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            "https://github.com/grafana/dataplane/edit/main/docusaurus/website",
        },
        theme: {
          customCss: require.resolve("./src/css/custom.css"),
        },
        blog: false,
      }),
    ],
  ],
  customFields: {
    rudderStackTracking: {
      url: 'https://rs.grafana.com',
      writeKey: '1sBAgwTlZ2K0zTzkM8YTWorZI00',
      configUrl: 'https://rsc.grafana.com',
      sdkUrl: 'https://rsdk.grafana.com',
    },
    canSpamUrl: 'https://grafana.com/canspam',
    gcomUrl: 'https://grafana.com/api',
    oneTrust: {
      enabled: true,
      scriptSrc: 'https://cdn.cookielaw.org/scripttemplates/otSDKStub.js',
      domainId: '019644f3-5dcf-741c-8b6d-42fb8feae57f-dev',
      analyticsGroupId: 'C0002', // OneTrust group ID for analytics consent
    },
  },
  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      navbar: {
        title: "Grafana Data Plane",
        logo: {
          alt: "Grafana Logo",
          src: "img/logo.svg",
        },
        items: [
          {
            href: devPortalHome,
            label: "Portal Home",
            position: "right",
            target: "_self",
          },
          {
            href: "https://www.github.com/grafana/dataplane",
            label: "GitHub",
            position: "right",
          },
        ],
      },
      footer: {
        style: "dark",
        links: [
          {
            title: "Docs",
            items: [
              {
                label: "Contract",
                to: "/",
              },
              {
                label: "Portal Home",
                href: devPortalHome,
                target: "_self",
              },
            ],
          },
          {
            title: "Tools & Examples",
            items: [
              {
                label: "Mock Data Source Plugin",
                href: "https://grafana.com/plugins/grafana-mock-datasource",
              },
              {
                label: "Example Data Frames (JSON)",
                href: "https://github.com/grafana/dataplane/tree/main/examples/data",
              },
              {
                label: "Go Testing/Example Library",
                href: "https://pkg.go.dev/github.com/grafana/dataplane/examples",
              },
              {
                label: "Go Dataplane Library",
                href: "https://pkg.go.dev/github.com/grafana/dataplane/sdata",
              },
            ],
          },
          {
            title: "Other Resources",
            items: [
              {
                label: "Go Plugin Data Package",
                href: "hhttps://pkg.go.dev/github.com/grafana/grafana-plugin-sdk-go/data",
              },
            ],
          },
          {
            title: "Community",
            items: [
              {
                label: "GitHub",
                href: "https://www.github.com/grafana/dataplane",
              },
              {
                label: "Github Issues",
                href: "https://www.github.com/grafana/dataplane/issues",
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} Grafana Labs. Built with Docusaurus.`,
      },
      prism: {
        theme: grafanaPrismTheme,
      },
      colorMode: {
        defaultMode: "dark",
        disableSwitch: true,
        respectPrefersColorScheme: false,
      },
    }),
};

module.exports = config;
