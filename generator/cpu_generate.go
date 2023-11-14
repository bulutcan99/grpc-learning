package generator

import (
	"grpc_fundamental/proto"
	"math/rand"
)

func NewCpu() *grpc_fundamental.Cpu {
	return &grpc_fundamental.Cpu{
		Brand:      RandomCpuBrand(),
		Name:       RandomCpuName(),
		NumCores:   RandomCpuCore(),
		NumThreads: RandomCpuCore() * 2,
		MaxFreq:    42,
		MinFreq:    21,
	}
}

func RandomCpuName() string {
	switch rand.Intn(3) {
	case 0:
		return "i7-8700K"
	case 1:
		return "i5-8600K"
	case 2:
		return "i3-8350K"
	default:
		return "i7-8700K"
	}
}

func RandomCpuBrand() string {
	switch rand.Intn(2) {
	case 0:
		return "Intel"
	case 1:
		return "AMD"
	default:
		return "Intel"
	}
}

func RandomCpuCore() uint32 {
	switch rand.Intn(3) {
	case 0:
		return 4
	case 1:
		return 6
	case 2:
		return 8
	default:
		return 8
	}
}
