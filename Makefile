fmt:
	gofmt -s -w -l `find . -name '*.go' -print`

swag:
	$(GOPATH)/bin/swag init -g cmd/hstreamdb-server/main.go

.PHONY: fmt, swag

