# go-tron-cli
go version [tron-cli](https://github.com/TRON-US/go-tron-cli)

## Install

### Use without build

1. Get binary from [release](https://github.com/TRON-US/go-tron-cli/releases)

2. There you go.

```troncli```

### Build from source

1. Get code

```
git clone https://github.com/TRON-US/go-tron-cli
``` 

2. Build

```
make build
```

3. It's there for you in ```/bin``` directory

## Usage

```
NAME:
   troncli - go version tron-cli

USAGE:
   troncli [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     init, i    init
     config, c  config
     run, c     run
     help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --typen value, --tn value  network type (default: "mainnet")
   --portf value, --pf value  fullnode port (default: "8090")
   --ports value, --ps value  solidity port (default: "8091")
   --help, -h                 show help
   --version, -v              print the version
```

### Set up node to sync with Mainnet

```
troncli init
```

```
troncli config
```

```
troncli run
```

### Set up node to sync with Privatenet

```
troncli init
```

```
troncli --typen private config
```

```
troncli run
```


## CHANGELOG

### v0.0.1

* Very basic version of tron-cli, yet can help you both set up a node sync to mainnet or set up private net node.

* Support java-tron 3.6.1

### v0.0.2

* Support mac, linux and windows distributions.
