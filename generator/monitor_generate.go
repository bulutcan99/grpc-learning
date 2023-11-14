package generator

import (
	"grpc_fundamental/proto"
)

func NewMonitor() *grpc_fundamental.Screen {
	return &grpc_fundamental.Screen{
		SizeInch: 27,
		Resolution: &grpc_fundamental.Screen_Resolution{
			Width:  1920,
			Height: 1080,
		},
		Panel: grpc_fundamental.Screen_IPS,
		Touch: true,
	}
}
