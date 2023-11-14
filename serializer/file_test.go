package serializer

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"grpc_fundamental/generator"
	"grpc_fundamental/proto"
	"testing"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	newLaptop := generator.NewLaptop()
	err := WriteProtobufToBinaryFile(newLaptop, binaryFile)
	require.NoError(t, err)
}

func TestReadProtobufFromBinaryFile(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	newLaptop := generator.NewLaptop()
	err := WriteProtobufToBinaryFile(newLaptop, binaryFile)
	require.NoError(t, err)

	laptop := &grpc_fundamental.Laptop{}
	err = ReadProtobufFromBinaryFile(laptop, binaryFile)
	require.NoError(t, err)
	require.True(t, proto.Equal(newLaptop, laptop))
}

func TestWriteProtobufToJSONFile(t *testing.T) {
	t.Parallel()
	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"
	laptopSample := generator.NewLaptop()
	err := WriteProtobufToBinaryFile(laptopSample, binaryFile)
	require.NoError(t, err)

	laptopSample2 := &grpc_fundamental.Laptop{}
	err = ReadProtobufFromBinaryFile(laptopSample2, binaryFile)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptopSample, laptopSample2))

	err = WriteProtobufToJSONFile(laptopSample2, jsonFile)
	require.NoError(t, err)

}
