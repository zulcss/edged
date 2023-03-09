FROM debian:testing

RUN apt-get update && \
    apt-get install -y golang build-essential ca-certificates
RUN mkdir /go 
WORKDIR /go
COPY . .
RUN make server && \
    cp bin/edged /usr/sbin/edged && \
    cp -rp etc/edged /etc
