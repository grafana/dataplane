---
id: contract
title: Grafana data structure 
description: Dataplanes and Dataframes.
keywords:
  - data-plane-contract
  - dataplane
  - dataframe
slug: /  
---

# Grafana data structure

Grafana supports a variety of different data sources, each with its own data model. To make this possible, Grafana consolidates the query results from each of these data sources into one unified data structure called a **data frame**. The **data plane** is built on top of the data frame and includes additional information on what the data frame holds. 

## Data frames overview  

A data frame is a data structure that consolidates the query results from your data sources, providing a common container in Grafana. 

:::caution

Query responses are often more than one single data frame. 

:::

The data frame a column-oriented table (the _fields_) with metadata (the _frame_) attached. Since data frame columns can have labels attached to them (`key=value`, `key2=val`...), it can hold Prometheus like responses as well. 

Each field in a data frame contains optional information about the values in the field, such as units, scaling, and so on. By adding field configurations to a data frame, Grafana can configure visualizations automatically. For example, you could configure Grafana to automatically set the unit provided by the data source.

## Data planes overview  

The data plane adds a property layer to the frame as metadata. It indicates the data frame _type_ (for example, a timeseries or a heatmap), which consists of a _kind_ (of data) and its _format_ (Prometheus-like, SQL-table-like).

:::tip

Data plane types are to data frames what TypeScript is to JavaScript.

:::

The data plane contract is a written set of rules that explain how producers of data (datasources, transformations) must form the frames, and how data consumers (like dashboards, alerting, and apps) can expect the data they receive to be like. In short, it describes the rules for valid and invalid schemas for each data frame type.

## Why use data planes?

Although the use of data planes is not enforced, the main objective of the data plane is to make Grafana more self-interoperable between data sources and features like dashboards and alerting. 

With data planes compatibility becomes about supporting data types and not specific features and data sources. For example, if data source produces type "A", and alerting and certain visualizations accept type "A", then that data source works with alerting and those visualizations.

### Benefits

Besides interoperability, using data planes has other benefits.

If you're a developer and data source author, you know what type of frames to output, and authors of features know what to expect for their input. This makes the platform scalable and development more efficient and less frustrating due to incompatibilities.

In general, using Grafana becomes more reliable, with everything working as expected. This helps to suggest what actions can be taken with the data. For example, if you're using a specific type, Grafana can suggest creating alert rules or certain visualizations in dashboards that work well with that type. Similarly, Grafana can suggest transformations that get you from the current type to another type support additional actions.

## What if I don't use data planes?

If you don't use data planes, consumers of data have to infer the type from the data returned, which has a few problems:

- Users are uncertain about how to write queries to work with different things.
- Error messages can become seemingly unrelated to what users are doing.
- Different features guess differently (for example, alerting vs. visualizations), making it hard for users and developers to know what to send.
- On the consumer side, guessing code becomes more convoluted over time as more exceptions are added for various data sources.

### What if my data source is schemaless and doesn't have kinds or types?

You can propose a new data plane type: They're designed to grow into maturity, not limit innovation.

Usually data sources have a drop down in the query UI to assert the query type, which appears as "format as". You can use this data source query information to produce a data plane-compatible type for the response.

While this may involve extra work for the user, defining the data plane type is easier at query time, since the data source knows more about the data that comes from the system behind the data source. 

## List of data sources that use data planes

As of October 2025, the following data sources send data plane data in at least some of their responses:

- Prometheus, including Amazon and Azure variants
- Loki
- Azure Monitor 
- Azure Data Explorer
- Bigquery
- Clickhouse
- Cloudlflare
- Databricks
- Influx
- MySQL
- New Relic
- Oracle 
- Postgres 
- Snowflake
- Victoria metrics

To see examples of data planes, refer to [data plane example data](https://github.com/grafana/dataplane/tree/main/examples/data) in GitHub.

## Read on

For more information on data frames refer to [Data frames](https://grafana.com/developers/plugin-tools/key-concepts/data-frames) in the Grafana Plugin Tools documentation.

To learn more about the data plane contract, see [Data Plane Contract - Technical Specification](./contract.md).
