package main

import (
	"context"
	"fmt"
	pb "lesson23/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type userService struct {
	pb.UnimplementedUserServiceServer
}

func (us *userService) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	if req.Id == "1" {
		return &pb.UserResponse{
			Id:   req.Id,
			Name: "Ivan",
		}, nil
	} else {
		return &pb.UserResponse{
			Id:   req.Id,
			Name: "Guest",
		}, nil
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &userService{})
	log.Println("gRPC started on :8080")
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println(err)
		return
	}

}
