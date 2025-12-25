package main

import (
	"context"
	"fmt"
	pb "lesson23/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.GetUser(ctx, &pb.UserRequest{Id: "2"})
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Printf("Получил данные: %s: %s", resp.Id, resp.Name)

}
