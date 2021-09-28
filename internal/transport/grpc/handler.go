package grpc

import (
	"context"
	"log"
	"net"

	usr "github.com/brandonmillerio/bmio_protos/nuboverflow/user/v1"
	"github.com/brandonmillerio/nuboverflow_users/internal/user"
	"google.golang.org/grpc"
)

type UserService interface {
	
	GetUserByID(ctx context.Context, id string) (user.User, error);
	CreateUser(ctx context.Context, user user.User) (user.User, error);
	UpdateUser(ctx context.Context, user user.User) (user.User, error);
	DeleteUser(ctx context.Context, id string) error;
}

type Handler struct {
	UserService UserService
	usr.UnimplementedUserServiceServer
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
	}

	grpcServer := grpc.NewServer()
	usr.RegisterUserServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("Feiled to serve: %s\n", err)
		return err
	}

	return nil
}

func (h Handler) GetUser(ctx context.Context, req *usr.GetUserRequest) (*usr.GetUserResponse, error) {
	return &usr.GetUserResponse{}, nil
}

func (h Handler) CreateUser(ctx context.Context, req *usr.CreateUserRequest) (*usr.CreateUserResponse, error) {
	return &usr.CreateUserResponse{}, nil
}

func (h Handler) UpdateUser(ctx context.Context, req *usr.UpdateUserRequest) (*usr.UpdateUserResponse, error) {
	return &usr.UpdateUserResponse{}, nil
}

func (h Handler) DeleteUser(ctx context.Context, req *usr.DeleteUserRequest) (*usr.DeleteUserResponse, error)  {
	return &usr.DeleteUserResponse{}, nil
}