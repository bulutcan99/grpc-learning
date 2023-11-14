package service

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_fundamental/generator"
	"grpc_fundamental/proto"
	"testing"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoId := generator.NewLaptop()
	laptopNoId.Id = ""

	laptopInvalidId := generator.NewLaptop()
	laptopInvalidId.Id = "invalid-uuid"

	laptopDuplicateId := generator.NewLaptop()
	storeDuplicateId := NewMapLaptopStore()
	err := storeDuplicateId.Save(laptopDuplicateId)
	require.Nil(t, err)

	testCases := []struct {
		name   string
		laptop *grpc_fundamental.Laptop
		store  LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: generator.NewLaptop(),
			store:  NewMapLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_with_id",
			laptop: generator.NewLaptop(),
			store:  NewMapLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopNoId,
			store:  NewMapLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "falied_invalid_id",
			laptop: laptopInvalidId,
			store:  NewMapLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failed_already_exist",
			laptop: laptopDuplicateId,
			store:  storeDuplicateId,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			req := &grpc_fundamental.CreateLaptopRequest{
				Laptop: testCase.laptop,
			}
			server := NewLaptopServer(testCase.store)
			res, err := server.CreateLaptop(req)
			if testCase.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(testCase.laptop.Id) > 0 {
					require.Equal(t, testCase.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				statusCode, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, testCase.code, statusCode.Code())
			}
		})
	}
}
