FROM ubuntu:focal as builder

RUN apt-get update && DEBIAN_FRONTEND="noninteractive" apt-get install -y \
      build-essential autoconf libtool libssl-dev pkg-config ca-certificates \
      golang-1.16-go && \
    ln -s /usr/lib/go-1.16/bin/go /usr/bin/go && \
    ln -s /usr/lib/go-1.16/bin/gofmt /usr/bin/gofmt

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
