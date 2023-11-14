package generator

func NewGpu() *pb.Gpu {
	return &pb.Gpu{
		Brand: "Nvidia",
		Name:  "GTX 1080",
		Memory: &pb.Memory{
			Value: 8,
			Unit:  pb.Memory_GIGABYTE,
		},
		MaxFreq: 2000,
		MinFreq: 1000,
	}
}
