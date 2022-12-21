package pb

import (
	"context"
	"fmt"
	"net"
	"universe-auth/internal/models"
	"universe-auth/internal/pb/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthWorker interface {
	Validate(ctx context.Context, token string) (bool, error)
	Login(ctx context.Context, email, password string) (*models.User, bool, error)
	Registration(ctx context.Context, email, usename, password string) (*models.User, error)
}

func NewServer(worker AuthWorker) {
	opt := grpc.Creds(insecure.NewCredentials())
	s := grpc.NewServer(opt)

	grpcAuthServer := auth.NewGrpcAuth(worker)
	auth.RegisterAuthServer(s, grpcAuthServer)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 50051))

	if err != nil {
		return
	}

	s.Serve(lis)
}
