package control

import (
	"context"
	"log/slog"
	"net"

	pb "github.com/tombuente/scara-proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedRobotServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ExecCommand(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	return &pb.CommandResponse{Response: req.Command}, nil
}

func Serve(listener net.Listener) {
	grpcServer := grpc.NewServer()
	pb.RegisterRobotServer(grpcServer, NewServer())

	slog.Info("Listening...")
	grpcServer.Serve(listener)
}
