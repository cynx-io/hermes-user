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
	pb.UnimplementedHermesUserServiceServer
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
	pb.RegisterHermesUserServiceServer(server, s)
	reflection.Register(server)

	logger.Infof("Starting gRPC server on %s", address)
	return server.Serve(lis)
}

func (s *Server) CheckUsername(ctx context.Context, req *pb.CheckUsernameRequest) (*pb.CheckUsernameResponse, error) {
	resp, err := s.userService.CheckUsername(ctx, req.Username)
	return serviceWrapper(resp, err)
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	resp, err := s.userService.GetUser(ctx, req.Username)
	return serviceWrapper(resp, err)
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	resp, err := s.userService.CreateUser(ctx, req.Username, req.Password)
	return serviceWrapper(resp, err)
}

func (s *Server) PaginateUsers(ctx context.Context, req *pb.PaginateRequest) (*pb.PaginateUsersResponse, error) {
	resp, err := s.userService.PaginateUsers(ctx, req)
	return serviceWrapper(resp, err)
}

func (s *Server) ValidatePassword(ctx context.Context, req *pb.ValidatePasswordRequest) (*pb.ValidatePasswordResponse, error) {
	resp, err := s.userService.ValidatePassword(ctx, req.Username, req.Password)
	return serviceWrapper(resp, err)
}
