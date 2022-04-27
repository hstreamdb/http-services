SWAG_PATH=$(shell go env GOPATH)/bin/swag

PACKAGE := github.com/hstreamdb/http-server

export GO_BUILD=GO111MODULE=on CGO_ENABLED=0 GOOS=$(GOOS) go build -ldflags '-s -w'

all: server adminCtl

fmt:
	gofmt -s -w -l `find . -name '*.go' -print`

swag:
	$(SWAG_PATH) init -g cmd/http-server/main.go

server:
	$(GO_BUILD) -o bin/http-server $(PACKAGE)/cmd/http-server

adminCtl:
	$(GO_BUILD) -o bin/adminCtl $(PACKAGE)/cmd/admin-client

.PHONY: fmt, swag, server, adminCtl, all

