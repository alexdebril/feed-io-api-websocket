FROM golang:buster AS builder

WORKDIR /build
COPY . /build

RUN go build

FROM debian:buster as runtime
COPY --from=builder /build/feed-io-api-websocket /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/feed-io-api-websocket"]
