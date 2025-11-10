// @ts-check
/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */
const sidebars = {
  doc: {
    Contract: [
      { id: "contract", label: "Intro", type: "doc" },
      { id: "data-structure", label: "Data structure", type: "doc" },
      { id: "dataframes", label: "Data frames", type: "doc" },
      { id: "contract-spec", label: "Data plane contract spec", type: "doc" },
      { id: "timeseries", label: "Timeseries", type: "doc" },
      { id: "numeric", label: "Numeric", type: "doc" },
      { id: "logs", label: "Logs", type: "doc" },
      { id: "heatmap", label: "Heatmap", type: "doc" },
    ],
  },
};

module.exports = sidebars;
