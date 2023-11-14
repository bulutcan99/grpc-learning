package generator

func NewMonitor() *pb.Screen {
	return &pb.Screen{
		SizeInch: 27,
		Resolution: &pb.Screen_Resolution{
			Width:  1920,
			Height: 1080,
		},
		Panel: pb.Screen_IPS,
		Touch: true,
	}
}
