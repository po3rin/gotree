# gotree

<img src="https://img.shields.io/badge/go-v1.11.1-blue.svg"/> [![CircleCI](https://circleci.com/gh/po3rin/gotree.svg?style=shield)](https://circleci.com/gh/po3rin/gotree) <img src="https://api.codeclimate.com/v1/badges/ab325a21df4463bb1c29/maintainability" /></a> [![GolangCI](https://golangci.com/badges/github.com/po3rin/gotree.svg)](https://golangci.com)

tree command to display go package document. This helps to understand the role of the package. hidden directory and File is not yet supported.

## Install

```
$ go get -u github.com/po3rin/gotree
```

## Usage

Specify the path as an argument.

```
$ gotree example
.
├── example1 ---( Package example1 for test. )
│   ├── a.go
│   ├── b.go
│   └── example11 ---( Package example11 for test. )
│       ├── a.go
│       ├── b.go
│       └── doc.go
├── example2
│   ├── a.txt
│   ├── b.txt
│   └── example21 ---( Package example21 for test. )
│       ├── a.go
│       ├── a_test.go
│       └── b.go
└── example3
    └── a.txt
```

## Flags

```
Usage:
  gotree [flags]

Flags:
      --color     color output (default true)
  -h, --help      help for gotree
      --version   version for gotree
```

## Test

following command.

```
$ go test ./...
```