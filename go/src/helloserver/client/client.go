package main

import(
	"context"
	"log"
	"time"
	"google.golang.org/grpc"
	"helloserver"
)

const address = "localhost:7000"

func main(){
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
                log.Fatalf("did not connect: %v", err)
        }
        defer connection.Close()

	client := helloserver.NewServerServiceClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

	response, err := client.PingServer(ctx, &helloserver.ClientRequest{ClientMessage : "Hello, I am client"})

	if err != nil {
                log.Fatalf("Could not ping server, error : %v", err)
        }

        log.Printf("Server response : %s", response.ServerMessage)
}
