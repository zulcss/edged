FROM debian:testing as build

RUN apt-get update && \
    apt-get install -y golang build-essential ca-certificates

FROM build as stx-client
RUN mkdir -p /go
ADD . /go
WORKDIR /go
RUN make client &&  \
   cp bin/stx /usr/bin
WORKDIR /
