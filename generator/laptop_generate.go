package generator

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math"
	"math/rand"
	pb "pb/proto"
	"time"
)

func NewLaptop() *pb.Laptop {
	return &pb.Laptop{
		Id:    uuid.New().String(),
		Brand: "Asus",
		Name:  "ROG",
		Cpu:   NewCpu(),
		Ram:   16,
		Gpu:   NewGpu(),
		Storage: []*pb.Storage{
			NewStorage(),
		},
		Screen: NewMonitor(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: rand.Float32() * 5,
		},
		Price: &pb.Laptop_PriceUsd{
			PriceUsd: uint32(int(math.Round(rand.Float64() * 3000))),
		},
		RelaseDate:    timestamppb.New(time.Now()),
		HasCamera:     true,
		HasMicrophone: true,
	}
}
