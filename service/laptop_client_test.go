package service

import (
	"google.golang.org/grpc"
	"grpc_fundamental/proto"
	"testing"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()
}

func startClientTestServer(t *testing.T) (*LaptopServer, error) {
	LaptopServer := NewLaptopServer(NewMapLaptopStore())
	grpcServer := grpc.NewServer()
	grpc_fundamental.RegisterLaptopServiceServer(grpcServer, LaptopServer)
}
