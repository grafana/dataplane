type NavCategory = "left" | "right";
type NavItem = { title: string; link: string; external?: boolean };

export const nav: Record<NavCategory, Array<NavItem>> = {
  left: [
    {
      title: "Contract",
      link: "/contract/contract",
      external: false,
    },
    {
      title: "Numeric",
      link: "/contract/numeric",
      external: false,
    },
    {
      title: "Timeseries",
      link: "/contract/timeseries",
      external: false,
    },
    {
      title: "Logs",
      link: "/contract/logs",
      external: false,
    },
    {
      title: "Heatmap",
      link: "/contract/heatmap",
      external: false,
    },
  ],
  right: [
    {
      title: "Grafana",
      link: "https://grafana.com",
      external: true,
    },
    {
      title: "Mock datasource plugin",
      link: "https://grafana.com/grafana/plugins/grafana-mock-datasource",
      external: true,
    },
  ],
};
