package grpc

import (
	"context"
	core "github.com/cynxees/cynx-core/proto/gen"
	grpccore "github.com/cynxees/cynx-core/src/grpc"
	"github.com/cynxees/cynx-core/src/logger"
	pb "github.com/cynxees/hermes-user/api/proto/gen/hermes"
	"github.com/cynxees/hermes-user/internal/module/usermodule"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

func (s *Server) Start(ctx context.Context, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterHermesUserServiceServer(server, s)
	reflection.Register(server)

	logger.Info(ctx, "Starting gRPC server on ", address)
	return server.Serve(lis)
}

func (s *Server) CheckUsername(ctx context.Context, req *pb.UsernameRequest) (resp *pb.CheckUsernameResponse, err error) {
	return grpccore.HandleGrpc(ctx, req, resp, s.userService.CheckUsername)
}

func (s *Server) GetUser(ctx context.Context, req *pb.UsernameRequest) (resp *pb.UserResponse, err error) {
	return grpccore.HandleGrpc(ctx, req, resp, s.userService.GetUser)
}

func (s *Server) CreateUser(ctx context.Context, req *pb.UsernamePasswordRequest) (resp *pb.UserResponse, err error) {
	return grpccore.HandleGrpc(ctx, req, resp, s.userService.CreateUser)
}

func (s *Server) PaginateUsers(ctx context.Context, req *pb.PaginateRequest) (resp *pb.PaginateUsersResponse, err error) {
	return grpccore.HandleGrpc(ctx, req, resp, s.userService.PaginateUsers)
}

func (s *Server) ValidatePassword(ctx context.Context, req *pb.UsernamePasswordRequest) (resp *pb.UserResponse, err error) {
	return grpccore.HandleGrpc(ctx, req, resp, s.userService.ValidatePassword)
}

func (s *Server) UpsertGuestUser(ctx context.Context, req *core.GenericRequest) (resp *pb.UserResponse, err error) {
	return grpccore.HandleGrpc(ctx, req, resp, s.userService.UpsertGuestUser)
}
