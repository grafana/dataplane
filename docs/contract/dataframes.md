# Data frames

A data frame is a collection of fields, where each field corresponds to a column. Each field, in turn, consists of a collection of values and metadata, such as the data type of those values. 

## Work with data frames

Refer to [Data frames](https://grafana.com/developers/plugin-tools/key-concepts/data-frames) in the Plugins developer documentation for an overview of data frames. 

For a guide to plugin development with data frames, refer to [Create data frames](https://grafana.com/developers/plugin-tools/how-to-guides/data-source-plugins/create-data-frames).

## Technical references

Data frames were introduced in Grafana 7.0, replacing the Time series and Table structures with a more generic data structure that can support a wider range of data types. The concept of data frame is borrowed from data analysis tools like the [R programming language](https://www.r-project.org) and [Pandas](https://pandas.pydata.org/). Other technical references are provided below.

### Apache Arrow

The data frame structure is inspired by, and uses the [Apache Arrow Project](https://arrow.apache.org/). Javascript Data frames use Arrow Tables as the underlying structure, and the backend Go code serializes its Frames in Arrow Tables for transmission.

### Javascript

The JavaScript implementation of data frames is in the [`/src/dataframe` folder](https://github.com/grafana/grafana/tree/main/packages/grafana-data/src/dataframe) and [`/src/types/dataframe.ts`](https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/types/dataFrame.ts) of the [`@grafana/data` package](https://github.com/grafana/grafana/tree/main/packages/grafana-data).

### Go

For documentation on the Go implementation of data frames, refer to the [github.com/grafana/grafana-plugin-sdk-go/data package](https://pkg.go.dev/github.com/grafana/grafana-plugin-sdk-go/data?tab=doc).






