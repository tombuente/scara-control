package control

import (
	"context"

	pb "github.com/tombuente/scara-proto"
)

type Server struct {
	pb.UnimplementedRobotServer

	service *Service
}

func NewServer(service *Service) *Server {
	s := Server{
		service: service,
	}

	return &s
}

func (s *Server) ExecCommand(ctx context.Context, req *pb.ExecCommandRequest) (*pb.ExecCommandResponse, error) {
	cmd := command{command: req.Command}
	id := s.service.addCommand(cmd)

	return &pb.ExecCommandResponse{Id: id}, nil
}
