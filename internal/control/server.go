package control

import (
	"context"
	"fmt"

	pb "github.com/tombuente/scara-proto"
)

type Server struct {
	pb.UnimplementedScaraServer

	service *Service
}

func NewServer(service *Service) *Server {
	s := Server{
		service: service,
	}

	return &s
}

func (s *Server) UploadProgram(ctx context.Context, req *pb.UploadProgramRequest) (*pb.Empty, error) {
	fmt.Println(req.Program)

	// cmd := command{command: req.Command}
	// id := s.service.addCommand(cmd)

	return &pb.Empty{}, nil
}
