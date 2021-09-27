package grpc

import (
	"context"
	"log"
	"net"

	usr "github.com/brandomillerio/nuboverflow_users/internal/user"
	user "github.com/brandonmillerio/bmio_protos/nuboverflow/user/v1"
	"google.golang.org/grpc"
)

type UserService interface {
	GetUserByID(ctx context.Context, id string) (usr.User, error);
	CreateUser(ctx context.Context, usr usr.User) (usr.User, error);
	UpdateUser(ctx context.Context, usr usr.User) (usr.User, error);
	DeleteUser(ctx context.Context, id string) error;
}

type Handler struct {
	UserService UserService
}

func New(userSerivice UserService) Handler {
	return Handler{
		UserService: userSerivice,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Print("Could not listen on port 50051")
		return err
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Print("Feiled to serve: %s\n", err)
		return err
	}

	return nil
}

