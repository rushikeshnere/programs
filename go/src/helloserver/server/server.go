package main

import (
	"net"
        "log"
	"context"

	"google.golang.org/grpc"
	"helloserver"
)

const (
	port = ":7000"	
)

type sampleServer struct {

} 

func (s *sampleServer) PingServer(ctx context.Context, req *helloserver.ClientRequest) (*helloserver.ServerResponse, error) {
	log.Printf("Client request : %s", req.ClientMessage)
	return &helloserver.ServerResponse{ServerMessage : "Hello, I am a server."}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
                log.Fatalf("Failed to listen on port: %v", err)
        }

	server := grpc.NewServer()
	helloserver.RegisterServerServiceServer(server, &sampleServer{})
	if err := server.Serve(listener); err != nil {
                log.Fatal("Failed to start server!", err)
        }
}
