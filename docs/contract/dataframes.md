# Data frames

A data frame is a collection of fields, where each field corresponds to a column. Each field, in turn, consists of a collection of values and metadata, such as the data type of those values. 

For example:

```ts
export interface Field<T = any, V = Vector<T>> {
  /**
   * Name of the field (column)
   */
  name: string;
  /**
   *  Field value type (string, number, and so on)
   */
  type: FieldType;
  /**
   *  Meta info about how field and how to display it
   */
  config: FieldConfig;

  /**
   * The raw field values
   * In Grafana 10, this accepts both simple arrays and the Vector interface
   * In Grafana 11, the Vector interface has been removed
   */
  values: V | T[];

  /**
   * When type === FieldType.Time, this can optionally store
   * the nanosecond-precison fractions as integers between
   * 0 and 999999.
   */
  nanos?: number[];

  labels?: Labels;

  /**
   * Cached values with appropriate display and id values
   */
  state?: FieldState | null;

  /**
   * Convert a value for display
   */
  display?: DisplayProcessor;

  /**
   * Get value data links with variables interpolated
   */
  getLinks?: (config: ValueLinkConfig) => Array<LinkModel<Field>>;
}
```

## Available data frame types

A data frame type consists of a **kind** of data (for example, time series or numeric) and the data **format** (for example, wide). 

The following data frame types are available:

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

## Work with data frames

Refer to [Data frames](https://grafana.com/developers/plugin-tools/key-concepts/data-frames) in the Plugins developer documentation for an overview of data frames. 

For a guide to plugin development with data frames, refer to [Create data frames](https://grafana.com/developers/plugin-tools/how-to-guides/data-source-plugins/create-data-frames).

## Technical references

Data frames were introduced in Grafana 7.0, replacing the Time series and Table structures with a more generic data structure that can support a wider range of data types. The concept of data frame is borrowed from data analysis tools like the [R programming language](https://www.r-project.org) and [Pandas](https://pandas.pydata.org/). Other technical references are provided below.

### Apache Arrow

The data frame structure is inspired by and uses the [Apache Arrow Project](https://arrow.apache.org/). JavaScript data frames use Arrow Tables as the underlying structure, and the backend Go code serializes its frames in Arrow Tables for transmission.

### JavaScript

You can find the JavaScript implementation of data frames in the [`/src/dataframe` folder](https://github.com/grafana/grafana/tree/main/packages/grafana-data/src/dataframe) and [`/src/types/dataframe.ts`](https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/types/dataFrame.ts) of the [`@grafana/data` package](https://github.com/grafana/grafana/tree/main/packages/grafana-data).

### Go

For documentation on the Go implementation of data frames, refer to the [Grafana SDK Go data package](https://pkg.go.dev/github.com/grafana/grafana-plugin-sdk-go/data?tab=doc).






