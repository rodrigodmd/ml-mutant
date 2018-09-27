# ML Mutant excercise [![Go Report Card](https://goreportcard.com/badge/github.com/rodrigodmd/ml-mutant)](https://goreportcard.com/report/github.com/rodrigodmd/ml-mutant) [![GoDoc](https://godoc.org/github.com/micro/go-api?status.svg)](https://godoc.org/github.com/rodrigodmd/ml-mutant) [![Build Status](https://travis-ci.org/rodrigodmd/ml-mutant.svg?branch=master)](https://travis-ci.org/rodrigodmd/ml-mutant) [![codecov](https://codecov.io/gh/rodrigodmd/ml-mutant/branch/master/graph/badge.svg)](https://codecov.io/gh/rodrigodmd/ml-mutant)

This is an example algorithm to resolve a mutant dna detection exercise

## Getting Started

To run the test you only need to [install go](https://golang.org/doc/install).

Then get the repository using go get:

    go get github.com/rodrigodmd/ml-mutant

Go to your project folder:

    cd $GOPATH/src/github.com/rodrigodmd/ml-mutant

## Run unit tests

To run the test you can use the following command:
```
make test
```
or

```
go test ./... -v -cover -race
```
