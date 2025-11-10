# Intro to the Grafana data structure

Grafana supports a variety of different data sources, each with its own data model. To make this possible, Grafana consolidates the query results from each of these data sources into one unified data structure called a **data frame**. The **data plane** adds a property layer to the data frame with information about the data frame type. 

Refer to [Grafana data structure](./data-structure.md) for an introduction to data frames and the [data plane layer](./contract-spec.md).



