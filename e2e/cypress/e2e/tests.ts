export const tests: {
  description: string;
  url: string;
  tablePanelChecks?: string[];
}[] = [
  {
    description:
      "numeric/long/v0.1/basic_valid/numeric-long_empty-two-item-names.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/long/v0.1/basic_valid/numeric-long_empty-two-item-names.json",
    tablePanelChecks: ["avgSlothCount", "avgSleepHoursPerSlothPerDay", "city"],
  },
  {
    description:
      "numeric/long/v0.1/basic_valid/numeric-long_four-items-by-name-and-dimension-two-labels.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/long/v0.1/basic_valid/numeric-long_four-items-by-name-and-dimension-two-labels.json",
    tablePanelChecks: [
      "avgSlothCount",
      "avgSleepHoursPerSlothPerDay",
      "city",
      "animal",
      "4",
      "23.5",
      "LGA",
      "cat",
    ],
  },
  {
    description:
      "numeric/long/v0.1/basic_valid/numeric-long_four-items-by-name-and-dimension.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/long/v0.1/basic_valid/numeric-long_four-items-by-name-and-dimension.json",
    tablePanelChecks: [
      "city",
      "avgSlothCount",
      "avgSleepHoursPerSlothPerDay",
      "LGA",
      "1",
      "23.5",
    ],
  },
  {
    description: "numeric/long/v0.1/basic_valid/numeric-long_no-data.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/long/v0.1/basic_valid/numeric-long_no-data.json",
    tablePanelChecks: ["No data"],
  },
  {
    description:
      "numeric/long/v0.1/basic_valid/numeric-long_two-items-by-dimension.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/long/v0.1/basic_valid/numeric-long_two-items-by-dimension.json",
    tablePanelChecks: ["avgSlothCount", "city", "4", "LGA", "7.50", "MIA"],
  },
  {
    description:
      "numeric/multi/v0.1/basic_valid/numeric-multi_four-items-by-dimension-name.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/multi/v0.1/basic_valid/numeric-multi_four-items-by-dimension-name.json",
    tablePanelChecks: ["avgSlothCount LGA", "4", `{city="LGA"}`],
  },
  {
    description: "numeric/multi/v0.1/basic_valid/numeric-multi_no-data.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/multi/v0.1/basic_valid/numeric-multi_no-data.json",
    tablePanelChecks: ["No data"],
  },
  {
    description:
      "numeric/multi/v0.1/basic_valid/numeric-multi_two-empty-items.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/multi/v0.1/basic_valid/numeric-multi_two-empty-items.json",
    tablePanelChecks: ["avgSlothCount LGA", "No data", `{city="LGA"}`],
  },
  {
    description:
      "numeric/multi/v0.1/basic_valid/numeric-multi_two-items-by-dimension-name-dif-name-dim.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/multi/v0.1/basic_valid/numeric-multi_two-items-by-dimension-name-dif-name-dim.json",
    tablePanelChecks: ["avgSlothCount LGA", "4", `{city="LGA"}`],
  },
  {
    description:
      "numeric/multi/v0.1/basic_valid/numeric-multi_two-items-by-dimension.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/multi/v0.1/basic_valid/numeric-multi_two-items-by-dimension.json",
    tablePanelChecks: ["avgSlothCount LGA", "4", `{city="LGA"}`],
  },
  {
    description:
      "numeric/multi/v0.1/extended_valid/numeric-multi_two-items-by-dimension-with-remainder-time.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/multi/v0.1/extended_valid/numeric-multi_two-items-by-dimension-with-remainder-time.json",
    tablePanelChecks: [
      "t",
      "avgSlothCount LGA",
      "2022-10-04 16:44:05",
      "4",
      `{city="LGA"}`,
    ],
  },
  {
    description:
      "numeric/wide/v0.1/basic_valid/numeric-wide_four-items-by-dimension-name.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/wide/v0.1/basic_valid/numeric-wide_four-items-by-dimension-name.json",
    tablePanelChecks: [
      "avgSlothCount LGA",
      "avgSlothCount MIA",
      "avgSleepHoursPerSlothPerDay LGA",
      "avgSleepHoursPerSlothPerDay MIA",
      "4",
      "7.50",
      "23.5",
      "23.2",
    ],
  },
  {
    description: "numeric/wide/v0.1/basic_valid/numeric-wide_no-data.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/wide/v0.1/basic_valid/numeric-wide_no-data.json",
    tablePanelChecks: ["No data"],
  },
  {
    description:
      "numeric/wide/v0.1/basic_valid/numeric-wide_two-empty-items.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/wide/v0.1/basic_valid/numeric-wide_two-empty-items.json",
    tablePanelChecks: ["avgSlothCount LGA", "avgSlothCount MIA"],
  },
  {
    description:
      "numeric/wide/v0.1/basic_valid/numeric-wide_two-items-by-dimension-name.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/wide/v0.1/basic_valid/numeric-wide_two-items-by-dimension-name.json",
    tablePanelChecks: [
      "avgSlothCount LGA",
      "avgSleepHoursPerSlothPerDay MIA",
      "4",
      "7.50",
    ],
  },
  {
    description:
      "numeric/wide/v0.1/basic_valid/numeric-wide_two-items-by-dimension.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/wide/v0.1/basic_valid/numeric-wide_two-items-by-dimension.json",
    tablePanelChecks: ["avgSlothCount LGA", "avgSlothCount MIA", "4", "7.50"],
  },
  {
    description:
      "numeric/wide/v0.1/extended_valid/numeric-wide_two-items-by-dimension-with-remainder-time.json",
    url: "https://raw.githubusercontent.com/grafana/dataplane/main/examples/data/numeric/wide/v0.1/extended_valid/numeric-wide_two-items-by-dimension-with-remainder-time.json",
    tablePanelChecks: [
      "t",
      "avgSlothCount LGA",
      "avgSlothCount MIA",
      "2022-10-04 16:44:05",
      "4",
      "7.50",
    ],
  },
];
