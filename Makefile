generate:
	protoc -I src/ --go_out=src/ src/protos/*.proto;

clean:
	rm src/grpc_fundamental/*.pb.go;

run:
	go run main.go
