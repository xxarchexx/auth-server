package e2e

import (
	"context"
	"testing"
	"universe-auth/internal/pb/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// дописать e2e тест
func TestRegistration(t *testing.T) {
	s := auth.NewGrpcAuth(worker)
	request := &auth.RegistrationRequest{Email: "test@mail.ru", Password: "12345", Username: "xxarchexx"}
	_, err := s.Registration(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}

}

func TestRegistratione2e(t *testing.T) {
	con, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer con.Close()
	client := auth.NewAuthClient(con)
	request := &auth.RegistrationRequest{Email: "test2@mail.ru", Password: "12345", Username: "xxarchexx"}
	_, err = client.Registration(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}

}
