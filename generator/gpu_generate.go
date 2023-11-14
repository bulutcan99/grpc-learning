package generator

import (
	"grpc_fundamental/proto"
)

func NewGpu() *grpc_fundamental.Gpu {
	return &grpc_fundamental.Gpu{
		Brand: "Nvidia",
		Name:  "GTX 1080",
		Memory: &grpc_fundamental.Memory{
			Value: 8,
			Unit:  grpc_fundamental.Memory_GIGABYTE,
		},
		MaxFreq: 2000,
		MinFreq: 1000,
	}
}
