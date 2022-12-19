# execserver

`execserver` is a command line application that executes arbitrary Go code.

NOT FOR PRODUCTION USE

## Dependencies

`execserver` depends on the `go` and `goimports` commands being available in the PATH.

## Install

```bash
go install github.com/mikerybka/execserver/cmd/execserver@latest
```

## Setup

1. Place your Go code in a Go Workspace.

## Usage

```bash
execserver <port> <auth_directory> <source_directory>
```

## API
