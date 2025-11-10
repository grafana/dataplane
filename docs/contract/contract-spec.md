# Data Plane Contract - Technical Specification

Grafana supports a variety of different data sources, each with its own data model. To make this possible, Grafana consolidates the query results from each of these data sources into one unified data structure called a **data frame**. The **data plane** adds a property layer to the data frame with information about the data frame type. Read [Grafana data structure](./dataplane-dataframes.md) for an introduction to data frames and the data plane layer.

## How is the data plane layer built?

The data plane layer indicates the data frame **type** (for example: time series data, numeric, or a histogram). In turn, the data frame type consists of a **kind** (of data) and the data **format** (Prometheus-like, SQL-table-like). 

For example, the `TimeSeriesWide` type consists of the kind "Time Series" and the format "Wide".

## Available data types

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

## Data sets 

A data type (kind+format) can have multiple items, forming a **set** of data items. For example, the numeric kind can have a set of numbers, or the time series kind can have a set of time series.

Each item of data in a set is uniquely identified by its **name** and its **dimensions**. Dimensions are facets of data (such as "location" or "host") with a corresponding value. For example, {"host"="a", "location"="new york"}.

In a data frame, dimensions are in either a field's label property or in string field.

### Properties of dimensional set-based kinds

- If multiple items have the same name, they need to have different dimensions (for example, labels) that uniquely identifies each item.
- The item name should appear in the `name` property of each value (numeric or bool typed) field, as any other label.
- A response can have different item names in the response. Note that Server Side Expressions (SSE) doesn't currently handle this option.

## Remainder data

Data is encoded into data frame(s), therefore all types are implemented as an array of `data.Frame`.

There can be data in data frame(s) that's not part of the data type's data. This extra data is the **remainder data** and is free to be used as convenient. What data becomes remainder data is dependent on and specified in the data type. Generally, it can be additional frames and/or additional fields of a certain field type.

:::caution
If you chose to use reminder data, libraries based on this contract must clearly delineate remainder data from data that is part of the type.
:::

## Possible responses

### Empty item response

If you retrieve one or more data items from a data source but an item has no values, that item is said to be an **"Empty value"**. In this case, the required data frame fields need to be present, although the fields themselves will have no values.

### "No Data" response

If a response has no data items, the response is a **"No Data"** response.

No data response can happen:

- If the entire set has no items.
- If the response has no frame and no error is returned.

If you have a response with no data, send a single frame (containing the data type, if applicable) and don't use any other fields on that frame.

### Invalid data response

If a data type is specified but the response doesn't follow the data type's rules, you'll get an error.

### Error responses

If a query returns an error, the error response is returned from outside the data frames using the `Error` and `Status` properties on [DataResponse](https://pkg.go.dev/github.com/grafana/grafana-plugin-sdk-go/backend#DataResponse). When an error is returned with the DataResponse, a single frame with no fields may be included as well, but it won't be considered **"No Data"** due to the error. This frame is included so that metadata, in particular a frame's `ExecutedQueryString`, can be returned with the error.

In a plugin backend, the call [`DataQueryHandler`](https://pkg.go.dev/github.com/grafana/grafana-plugin-sdk-go/backend#QueryDataHandler) can return an error. Use this option only when the entire request (all queries) fail.

### Responses with multiple data types

:::caution
Multiple data type responeses are not supported at the moment.
:::

If you don't use multi-type responses, you'll get the first data type that matches what you're querying for.

Although not supported, if you need to use responses with multiple data types (within a `RefID`), the following applies:

- Responses might not work as expected.
- Use only one format per data type within a response. For example, you may use TimeSeriesWide and NumericLong, but do not mix TimeSeriesWide and TimeSeriesLong.
- Derive the borders between the types from adjacent frames (within the array of frames) that share the same data type.

## Versioning

:::important
The data plane contract needs to be as stable as possible. 
:::

Versioning recommendations:

- Use contract versions only for changes impacting overarching concepts such as error handling or multi-data type responses. In other words, when version `1.0` is reached, limit changes to enhancements before working on version `2.0`.
- The addition of new data types, or the modification of data types should not impact the contract version.

### Data type versions

Give each data type a version in major/minor form (x.x). The version is located in the `Frame.Meta.TypeVersion` property.

- Version `0.0` means the data type is either pre contract, or in very early development.
- Version `0.x` means the data type is well defined in the contract, but may change based on things learned for wider usage.
- Version `1.0` should be a stable data type, and should have no changes from the previous version.
- Minor version changes beyond `1.0` must be backward compatible for data reader implementations. They also must be forward compatible with other `1.x` versions for readers as well (but without enhancement support).


