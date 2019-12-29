package main

import(
        "context"
        "log"
        "time"
        "google.golang.org/grpc"
        "studentservice"
)

const address = "localhost:7000"

func main(){
        connection, err := grpc.Dial(address, grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %v", err)
        }
        defer connection.Close()

        client := studentservice.NewStudentServiceClient(connection)
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

        response, err := client.GetStudent(ctx, &studentservice.GetStudentRequest{StudentId : 13})

        if err != nil {
                log.Fatalf("Could not get student details, error : %v", err)
        }

        log.Printf("Server response : %v", response)
}
