SWAG_PATH=$(shell go env GOPATH)/bin/swag

PACKAGE := github.com/hstreamdb/http-server

export GO_BUILD=GO111MODULE=on CGO_ENABLED=0 GOOS=$(GOOS) go build -ldflags '-s -w'

fmt:
	gofmt -s -w -l `find . -name '*.go' -print`

swag:
	$(SWAG_PATH) init -g cmd/http-server/main.go

server:
	$(GO_BUILD) -o bin/http-server $(PACKAGE)/cmd/http-server

adminCli:
	$(GO_BUILD) -o bin/admin-cli $(PACKAGE)/cmd/admin-client

runServer:
	bin/http-server

runAdmin:
	bin/admin-cli

.PHONY: fmt, swag, server, adminCli, runServer, runAdmin

