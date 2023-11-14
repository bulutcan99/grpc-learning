package generator

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc_fundamental/proto"
	"math"
	"math/rand"
	"time"
)

func NewLaptop() *grpc_fundamental.Laptop {
	return &grpc_fundamental.Laptop{
		Id:    uuid.New().String(),
		Brand: "Asus",
		Name:  "ROG",
		Cpu:   NewCpu(),
		Ram:   16,
		Gpu:   NewGpu(),
		Storage: []*grpc_fundamental.Storage{
			NewStorage(),
		},
		Screen: NewMonitor(),
		Weight: &grpc_fundamental.Laptop_WeightKg{
			WeightKg: rand.Float32() * 5,
		},
		Price: &grpc_fundamental.Laptop_PriceUsd{
			PriceUsd: uint32(int(math.Round(rand.Float64() * 3000))),
		},
		RelaseDate:    timestamppb.New(time.Now()),
		HasCamera:     true,
		HasMicrophone: true,
	}
}
