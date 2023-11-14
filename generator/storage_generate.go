package generator

func NewStorage() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: 500,
			Unit:  pb.Memory_GIGABYTE,
		},
	}
}
