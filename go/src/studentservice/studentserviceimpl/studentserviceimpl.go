package main

import (
	"context"
	"net"
	"log"

	"google.golang.org/grpc"	
	"studentservice"	
)


const (
	port = ":7000"
)

type student_service struct {

}

func (s *student_service) GetStudent(ctx context.Context, req *studentservice.GetStudentRequest) (*studentservice.StudentData, error) {
	student_data := &studentservice.StudentData{StudentId : 1, StudentFirstName : "Rushikesh", StudentSecondName : "Nere"}

        return student_data, nil
}

func main() {
	
	listener, err := net.Listen("tcp", port)
	if err != nil {
                log.Fatalf("Failed to listen on port: %v", err)
        }

	server := grpc.NewServer()
	studentservice.RegisterStudentServiceServer(server, &student_service{})
	if err := server.Serve(listener); err != nil {
                log.Fatal("Failed to start server!", err)
        }
}
