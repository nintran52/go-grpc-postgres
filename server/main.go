package main

import (
	"context"
	"log"
	"net"

	pb "github.com/nintran52/go-grpc-postgres/proto"
	"google.golang.org/grpc"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, req *pb.NewTodo) (*pb.Todo, error) {
	log.Printf("received: %v", req.GetName())
	todo := &pb.Todo{
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}

	return todo, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("tcp connection failed: %v", err)
	}
	log.Printf("listening at %v", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &TodoServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("grpc server failed: %v", err)
	}
}
