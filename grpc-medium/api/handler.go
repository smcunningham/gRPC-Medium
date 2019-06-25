package api

import (
	context "context"
	"log"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping rquest
func (s Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage,
	error) {
	log.Printf("Handler: Received message from client: %s\n", in.Greeting)
	return &PingMessage{Greeting: "Handler: This is the response sent from the handler {bar}"}, nil
}
