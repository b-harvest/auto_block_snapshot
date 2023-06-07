# auto_block_snapshot

auto_block_snapshot is a tool designed to automate the process of snapshotting, pruning, and uploading data to an AWS S3 bucket from a blockchain full node built with Cosmos-SDK and Tendermint.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)

## Installation
This project is written in Go. Make sure you have Go installed on your machine.

```
git clone https://github.com/b-harvest/auto_block_snapshot.git
cd auto_block_snapshot
go build .
./auto_block_snapshot
or
./auto_block_snapshot -c /path/to/config.toml
```


## Configuration
The configuration of this tool can be done using a `config.toml` file. An example configuration can be found in `config.toml`.

The `config.toml` should contain the following fields:

- `fullnode.path`: path to the full node binary
- `aws.region`: AWS region
- `aws.bucket`: AWS S3 bucket name
