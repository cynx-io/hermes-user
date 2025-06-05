package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hermes/api/proto/user"
	"hermes/internal/module/usermodule"
	"hermes/internal/pkg/logger"
	"net"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	userService *usermodule.UserService
}

func NewServer(userService *usermodule.UserService) *Server {
	return &Server{
		userService: userService,
	}
}

func (s *Server) Start(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, s)
	reflection.Register(server)

	logger.Infof("Starting gRPC server on %s", address)
	return server.Serve(lis)
}

func (s *Server) CheckUsername(ctx context.Context, req *pb.CheckUsernameRequest) (*pb.CheckUsernameResponse, error) {
	return s.userService.CheckUsername(ctx, req.Username)
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.userService.GetUser(ctx, req.Username)
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.userService.CreateUser(ctx, req.Username, req.Password)
}

func (s *Server) PaginateUsers(ctx context.Context, req *pb.PaginateRequest) (*pb.PaginateUsersResponse, error) {
	return s.userService.PaginateUsers(ctx, req)
}
