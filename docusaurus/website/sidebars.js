// @ts-check
/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */
const sidebars = {
  doc: {
    Contract: [
      { id: "dataplane-dataframes", label: "Dataplanes and Dataframes", type: "doc" },
      { id: "contract", label: "Data Plane Contract Tech Spec", type: "doc" },
      { id: "timeseries", label: "Timeseries", type: "doc" },
      { id: "numeric", label: "Numeric", type: "doc" },
      { id: "logs", label: "Logs", type: "doc" },
      { id: "heatmap", label: "Heatmap", type: "doc" },
    ],
  },
};

module.exports = sidebars;
