# Intro to the Grafana data structure

Grafana supports a variety of different data sources, each with its own data model. To make this possible, Grafana consolidates the query results from each of these data sources into one unified data structure called a **data frame**. The **data plane** adds a property layer to the data frame with information about the data frame type. 

## Learn more

Read [Grafana data structure](./dataplane-dataframes.md) for an introduction to data frames and the [data plane layer](./contract-spec.md).

The following types are available:

- [Time series](./timeseries.md)
  - [Wide](./timeseries.md#time-series-wide-format-timeserieswide)
  - [Long](./timeseries.md#time-series-long-format-timeserieslong-sql-like)
  - [Multi](./timeseries.md#time-series-multi-format-timeseriesmulti)
- [Numeric](./numeric.md)
  - [Wide](./numeric.md#numeric-wide-format-numericwide)
  - [Multi](./numeric.md#numeric-multi-format-numericmulti)
  - [Long](./numeric.md#numeric-many-format-numericlong)
- [Heatmap](./heatmap.md)
  - [Rows](./heatmap.md#heatmap-rows-heatmaprows)
  - [Cells](./heatmap.md#heatmap-cells-heatmapcells)
- [Logs](./logs.md)
  - [LogLines](./logs.md#loglines)

