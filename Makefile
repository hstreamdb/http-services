fmt:
	gofmt -s -w -l `find . -name '*.go' -print`

.PHONY: fmt

