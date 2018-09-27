GOPATH:=$(shell go env GOPATH)

.PHONY: test

test:
	go test ./... -v -cover -race