# masterstat-cli [![build](https://github.com/vikpe/masterstat-cli/actions/workflows/build.yml/badge.svg)](https://github.com/vikpe/masterstat-cli/actions/workflows/build.yml)  [![codecov](https://codecov.io/gh/vikpe/masterstat-cli/branch/main/graph/badge.svg)](https://codecov.io/gh/vikpe/masterstat-cli) [![Go Report Card](https://goreportcard.com/badge/github.com/vikpe/masterstat-cli)](https://goreportcard.com/report/github.com/vikpe/masterstat-cli)

```shell
Fetch server addresses from QuakeWorld master servers.

  Usage:   masterstat [<address> ...]
Example:   masterstat master.quakeworld.nu:27000 qwmaster.ocrana.de:27000
```

## Download

See [releases](https://github.com/vikpe/masterstat-cli/releases) for downloads.

## Usage

### Single master server

```shell
masterstat master.quakeworld.nu:27000
```

```
193.200.16.105:28502
193.200.16.105:28504
193.200.16.105:28000
193.200.16.105:30000
91.211.246.220:28000
[...]
```

### Multiple master servers

**Note**: Returns unique server addresses (no duplicates).

```shell
masterstat master.quakeworld.nu:27000 qwmaster.ocrana.de:27000
```

```
193.200.16.105:28502
193.200.16.105:28504
193.200.16.105:28000
193.200.16.105:30000
91.211.246.220:28000
[...]
```

## Build from source

```shell
git clone git@github.com:vikpe/masterstat-cli.git
cd masterstat-cli
go build
```

## See also

* [masterstat](https://github.com/vikpe/masterstat)
* [serverstat](https://github.com/vikpe/serverstat)
* [serverstat-cli](https://github.com/vikpe/serverstat-cli)
