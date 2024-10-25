FROM golang:1.23.2 AS build

ARG VERSION
ARG DEBUG

COPY . /go/src/goapp
WORKDIR /go/src/goapp

RUN make

FROM debian:12-slim

RUN \
    apt-get -y update && \
    apt-get -y upgrade && \
    apt-get -y install ca-certificates iputils-ping net-tools netcat-traditional procps tzdata && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=build /go/src/goapp/bin /goapp

ENTRYPOINT ["/goapp/server"]