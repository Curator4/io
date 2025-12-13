package grpc

import (
	"errors"

	"github.com/curator4/io/backend/internal/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// toGRPCError maps core errors to appropriate gRPC status errors
func toGRPCError(err error) error {
	if err == nil {
		return nil
	}

	// Check for sentinel errors
	switch {
	case errors.Is(err, core.ErrNoConfigsFound):
		return status.Error(codes.FailedPrecondition, "system not configured: no AI configs found")
	case errors.Is(err, core.ErrProviderNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, core.ErrLLMUnavailable):
		return status.Error(codes.Unavailable, err.Error())
	default:
		// All other errors are internal server errors
		return status.Error(codes.Internal, err.Error())
	}
}
