.PHONY: server client clean docker agent proto

clean:
	rm -rf bin

server:
	go build -o bin/edged apiserver/main.go

client:
	go build -o bin/stx client/main.go

docker:
	docker build -t stx-client .

agent:
	go build -o bin/agent agent/main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/agent.proto
