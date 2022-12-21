package auth

import (
	context "context"
	"universe-auth/internal/errors"
	"universe-auth/internal/models"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"
)

type AuthWorker interface {
	Validate(ctx context.Context, token string) (bool, error)
	Login(ctx context.Context, email, password string) (*models.User, bool, error)
	Registration(ctx context.Context, email, username, password string) (*models.User, error)
}

type Auth struct {
	UnimplementedAuthServer
	AuthManager AuthWorker
}

func NewGrpcAuth(worker AuthWorker) *Auth {
	server := &Auth{
		AuthManager: worker,
	}
	return server
}

func (s *Auth) Validate(ctx context.Context, req *ValidateRequest) (*ValidateResponse, error) {
	res, err := s.AuthManager.Validate(ctx, req.Token)
	if err != nil || !res {
		return nil, err
	}

	return &ValidateResponse{IsValid: true}, nil
}

func (s *Auth) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {

	res, isLoggined, err := s.AuthManager.Login(ctx, req.Email, req.Password)

	if err != nil {
		return nil, err
	}

	if !isLoggined {
		response := LoginResponse{}
		return &response, nil
	}

	response := LoginResponse{UserName: res.Username, Email: res.Email, Token: res.Token}
	return &response, nil
}

func (s *Auth) Registration(ctx context.Context, req *RegistrationRequest) (*RegistrationResponse, error) {
	if req.Username == "" || req.Email == "" || req.Password == "" {
		err := status.Error(codes.InvalidArgument, "invalid arguments")
		return nil, err
	}

	res, err := s.AuthManager.Registration(ctx, req.Email, req.Username, req.Password)

	if err != nil {
		return nil, errors.HandleGrpcError(err)
	}

	response := RegistrationResponse{UserName: res.Username, Email: res.Email, Token: ""}
	return &response, nil
}
