package errors

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleGrpcError(err error) error {
	if errors.Is(err, AlreadyExists) {
		return status.Error(codes.AlreadyExists, err.Error())
	}

	if errors.Is(err, InvalidArgument) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}
