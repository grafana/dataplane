<div align="center">
  <img
    src="./docs/img/logo.svg"
    alt="Grafana Logo"
    width="100px"
    padding="40px"
  />
  <h1>Grafana Data Plane</h1>
  <p>Grafana data plane tools and docs</p>
</div>
<div align="center">
  <a href="https://github.com/grafana/dataplane/actions/workflows/ci.yml"
    ><img
      src="https://github.com/grafana/dataplane/actions/workflows/ci.yml/badge.svg"
      alt="Tests & builds status" /></a
  >
  <br />
  <br />
</div>

This is a monorepo of Grafana dataplane tools and docs

## Docs
- [Data Plane Contract - Technical Specification](https://grafana.github.io/dataplane/contract/)

## backend packages

- [sdata](./sdata/) (**experimental** Structural way of building typed dataframes)
- [examples](./examples/) (Examples of dataplane typed dataframes in json files, and a go library for using them in tests)
