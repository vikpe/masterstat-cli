# masterstat-cli [![build](https://github.com/vikpe/masterstat-cli/actions/workflows/build.yml/badge.svg)](https://github.com/vikpe/masterstat-cli/actions/workflows/build.yml)

```shell
Fetch server addresses from QuakeWorld master servers.

  Usage:   masterstat [<address> ...]
Example:   masterstat master.quakeworld.nu:27000 qwmaster.ocrana.de:27000
```

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
