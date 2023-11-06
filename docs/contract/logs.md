# Logs

Status: EARLY Draft/Proposal

## LogLines

Version: 0.0

### Properties and field requirements

- **Time field** - _required_
  - The first time field with the name `timestamp` is the time field.
  - it must be non nullable
- **Body field** - _required_
  - The first string field with the name `body` is the body field.
  - it must be non nullable
- **Severity field** - _optional_
  - The first string field with the name `severity` is the severity field.
  - Represents the severity/level of the log line
  - If no severity field is found, consumers/client will decide the log level. Example: logs panel will try to parse the message field and determine the log level
  - Log level can be one of the values specified in the docs [here](https://grafana.com/docs/grafana/latest/explore/logs-integration/)
- **ID field** - _optional_
  - The first string field with the name `id` is the id field.
  - Unique identifier of the log line
- **Labels field** - _optional_
  - The first field with the name `labels` is the labels field.
  - This field represent additional labels of the log line.
  - Field type must be json raw message type. Example value: `{}`, `{"hello":"world", "foo": 123.45, "bar" :["yellow","red"], "baz" : { "name": "alice" }}`
    - Value should be represented with `Record<string,any>` type in javascript.

Any other field is ignored by logs visualisation.

### Example

Following is an example of a logs frame in go

```go
data.NewFrame(
    "logs",
    data.NewField("timestamp", nil, []time.Time{time.UnixMilli(1645030244810), time.UnixMilli(1645030247027), time.UnixMilli(1645030247027)}),
    data.NewField("body", nil, []string{"message one", "message two", "message three"}),
    data.NewField("severity", nil, []string{"critical", "error", "warning"}),
    data.NewField("id", nil, []string{"xxx-001", "xyz-002", "111-003"}),
    data.NewField("labels", nil, []json.RawMessage{[]byte(`{}`), []byte(`{"hello":"world"}`), []byte(`{"hello":"world", "foo": 123.45, "bar" :["yellow","red"], "baz" : { "name": "alice" }}`)}),
)
```

the same can be represented as

| Name: timestamp <br/> Type: []time.Time | Name: body <br/> Type: []string | Name: severity <br/> Type: []\*string | Name: id <br/> Type: []\*string | Name: labels <br/> Type: []json.RawMessage                                         |
| --------------------------------------- | ------------------------------- | ------------------------------------- | ------------------------------- | -------------------------------------------------------------------------------------- |
| 2022-02-16 16:50:44.810 +0000 GMT       | message one                     | critical                              | xxx-001                         | {}                                                                                     |
| 2022-02-16 16:50:47.027 +0000 GMT       | message two                     | error                                 | xyz-002                         | {"hello":"world"}                                                                      |
| 2022-02-16 16:50:47.027 +0000 GMT       | message three                   | warning                               | 111-003                         | {"hello":"world", "foo": 123.45, "bar" :["yellow","red"], "baz" : { "name": "alice" }} |

### Meta data requirements

- Frame type must be set to `FrameTypeLogLines`/`log-lines`
- Frame meta can optionally specify `preferredVisualisationType:logs` as meta data. Without this property, explore page will be rendering the logs data as table instead in logs view

### Invalid cases

- Frame without time field
- Frame without string field
- Frame with field name "tsNs" where the type of the "tsNs" field is not number.

## Useful links

- [OTel Logs Data Model](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/data-model.md)
- [OTel Logs Level](https://docs.google.com/document/d/1WQDz1jF0yKBXe3OibXWfy3g6lor9SvjZ4xT-8uuDCiA/edit#)
- [Javascript high resolution timestamp](https://www.w3.org/TR/hr-time/)
