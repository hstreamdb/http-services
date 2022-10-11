FROM golang:1.19 as builder

COPY . /srv

RUN cd /srv && \
    export GO111MODULE=on CGO_ENABLED=0 GOOS=$GOOS && \
    go build -ldflags '-s -w' -v \
        -o /root/.local/bin/hstream-http-server \
        ./cmd/http-server && \
    rm -rf /srv

# -----------------------------------------------------------------------------

FROM ubuntu:focal

COPY --from=builder /root/.local/bin/hstream-http-server /usr/local/bin/
