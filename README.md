# gofilehead

_The gofilehead module matches a provided bytestream against the provided hexsignature._

## Table of contents

1. [Introduction](#introduction)
1. [Getting started](#getting-started)
    1. [Prerequisites](#prerequisites)
    1. [Installation](#installation)
1. [How to use](#how-to-use)
1. [License](/LICENSE)

## Introduction

This module matches a byte array, for example from a file or a tcp stream, against the provided hex signature.

If the result is true the by array matches the hexsignature file type.

A file with hex signatures can be found [here](./FileHeader.csv).

## Getting started

### Prerequisites

- Golang installed
- Operatingsystem: Linux or Windows, not tested on mac
- IDE like VS Code, if you want to contribute or change the code

### Installation

The recommended way to use this module is using the go cli.

    go get github.com/IT-Administrators/gofilehead

## How to use

```Go
// Read the example file.
cont := readFile("./examples/Gopher.png")
// Hexsignature of .png files.
gofilehead.MatchFileHead(cont, "89504E470D0A1A0A")

// Result:
// true
```

## License

[MIT](./LICENSE)