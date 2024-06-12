package main

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/tombuente/scara-control/internal/control"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 7777))
	if err != nil {
		slog.Error("Unable to create listener", "error", err)
	}

	control.Serve(listener)
}
