package grpc

import (
	"context"
	grpccore "github.com/cynx-io/cynx-core/src/grpc"
	"github.com/cynx-io/cynx-core/src/logger"
	pb "github.com/cynx-io/hermes-user/api/proto/gen/hermes"
	"github.com/cynx-io/hermes-user/internal/usecase/userusecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	pb.UnimplementedHermesUserServiceServer
	UseCase *userusecase.UseCase
}

func NewServer(useCase *userusecase.UseCase) *Server {
	return &Server{
		UseCase: useCase,
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

func (s *Server) UpsertUser(ctx context.Context, req *pb.UpsertUserRequest) (*pb.UserResponse, error) {
	var resp pb.UserResponse
	return grpccore.HandleGrpc(ctx, req, &resp, s.UseCase.UpsertUser)
}
