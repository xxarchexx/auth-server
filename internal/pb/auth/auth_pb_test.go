package auth

import (
	"context"
	"log"
	"net"
	"testing"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func TestLoginTest(t *testing.T) {
	s := Auth{}
	req := LoginRequest{Password: "123", Email: "ss@mail.ru"}
	_, err := s.Login(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRequestration(t *testing.T) {
	s := Auth{}
	request := &RegistrationRequest{Email: "ss@mail.ru", Password: "12345", Username: "xxarchexx"}
	_, err := s.Registration(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoginWithNetowrk(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := NewAuthClient(conn)
	resp, err := client.Login(ctx, &LoginRequest{Email: "ss@mail.ru", Password: "123"})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
}

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	RegisterAuthServer(s, &Auth{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
