SWAG_PATH=$(shell go env GOPATH)/bin/swag

fmt:
	gofmt -s -w -l `find . -name '*.go' -print`

swag:
	$(SWAG_PATH) init -g cmd/http-server/main.go

build:
	go build -o bin/http-server cmd/http-server/main.go

run: 
	bin/http-server

.PHONY: fmt, swag, build, run

