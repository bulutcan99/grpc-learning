package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	pb "pb/proto"
)

type LaptopServer struct {
	Store LaptopStore
	pb.UnimplementedLaptopServiceServer
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{
		Store: store,
	}
}

func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	fmt.Println("Received a create-laptop request with name: ", laptop.Name)
	if len(laptop.Id) == 0 {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, fmt.Errorf("cannot generate new id: %w", err)
		}

		laptop.Id = id.String()
	}
	err := server.Store.Save(laptop)
	if err != nil {
		if errors.Is(err, ErrAlreadyExist) {
			return nil, fmt.Errorf("laptop with id %s already exist: %w", laptop.Id, err)
		}
		return nil, fmt.Errorf("cannot save laptop to the store: %s", err)
	}

	fmt.Println("Saved laptop with id: ", laptop.Id)
	return &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}, nil
}
