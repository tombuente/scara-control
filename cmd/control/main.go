package main

import (
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"

	"github.com/tombuente/scara-control/internal/control"
	pb "github.com/tombuente/scara-proto"
)

func main() {
	controlService := control.NewService()

	serveRPC(controlService, "tcp", fmt.Sprintf("localhost:%d", 5000))
}

func serveRPC(service *control.Service, network string, address string) error {
	listener, err := net.Listen(network, address)
	if err != nil {
		return fmt.Errorf("unable to create listener: %w", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterScaraServer(grpcServer, control.NewServer(service))

	slog.Info("Listening...")
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("unable to serve grpc server: %w", err)
	}

	return nil
}
