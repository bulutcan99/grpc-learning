generate:
	protoc -I. --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/laptop.proto

clean:
	rm proto/*.pb.go;

run:
	go run main.go

test:
	go test -cover -race ./...