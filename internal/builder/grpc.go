package builder

import (
	grpc2 "github.com/tweeker88/auth-oleg/internal/transport/grpc"
	"github.com/tweeker88/auth-oleg/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func BuildServer() *grpc.Server {
	srv := grpc.NewServer()
	reflection.Register(srv)

	my := grpc2.NewServer()

	user_v1.RegisterUserV1Server(srv, my)

	return srv
}