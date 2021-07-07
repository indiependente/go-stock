[![Go Report Card](https://goreportcard.com/badge/github.com/indiependente/go-stock)](https://goreportcard.com/report/github.com/indiependente/go-stock)
# go-stock
A stock price viewer written in Go.

## Install
`go install github.com/indiependente/go-stock/cmd/go-stock`

## Usage
```bash
usage: go-stock [<flags>] [<symbol>]

Flags:
      --help                 Show context-sensitive help (also try --help-long and --help-man).
  -c, --config="config.yml"  Configuration file location
      --setup                Configure the tool will require you providing a few information

Args:
  [<symbol>]  Stock symbol
```

## Setup
The `--setup` flag will launch the setup process.
It will ask for which service you want to use among the ones supported by the tool.
And it will ask for its API key.

The configuration will then be saved in YAML format in a file in a location of your choice.

