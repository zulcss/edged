.PHONY: server client clean docker

clean:
	rm -rf bin

server:
	go build -o bin/edged apiserver/main.go

client:
	go build -o bin/stx client/main.go

docker:
	docker build -t stx-client .
