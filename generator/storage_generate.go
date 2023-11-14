package generator

import (
	"grpc_fundamental/proto"
)

func NewStorage() *grpc_fundamental.Storage {
	return &grpc_fundamental.Storage{
		Driver: grpc_fundamental.Storage_SSD,
		Memory: &grpc_fundamental.Memory{
			Value: 500,
			Unit:  grpc_fundamental.Memory_GIGABYTE,
		},
	}
}
