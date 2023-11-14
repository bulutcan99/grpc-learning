package service

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"pb/generator"
	pb "pb/proto"
	"testing"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopServer, serverAddress := startTestServer(t)
	laptopClient := newLaptopClient(t, serverAddress)

	laptop := generator.NewLaptop()
	expectedLaptopId := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedLaptopId, res.Id)

}

func startTestServer(t *testing.T) (*LaptopServer, string) {
	LaptopServer := NewLaptopServer(NewMapLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, LaptopServer)
	listener, err := net.Listen("tcp", ":8081")
	require.NoError(t, err)
	go grpcServer.Serve(listener)
	return LaptopServer, listener.Addr().String()
}

func newLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}
