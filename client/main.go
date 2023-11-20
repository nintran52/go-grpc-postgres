package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nintran52/go-grpc-postgres/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("problem with the server: %v", err)
	}
	defer conn.Close()

	c := pb.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	newTodos := []*pb.NewTodo{
		{Name: "Buy pizza", Description: "Make sure to get a spicy one.", Done: false},
		{Name: "Have fun", Description: "Eat and enjoy.", Done: false},
	}

	for _, newTodo := range newTodos {
		res, err := c.CreateTodo(ctx, &pb.NewTodo{Name: newTodo.Name, Description: newTodo.Description, Done: newTodo.Done})
		if err != nil {
			log.Fatalf("could not create todo: %v", err)
		}

		log.Printf(`
			ID: %s
			Name: %s
			Description: %s
			Done: %v
		`, res.GetId(), res.GetName(), res.GetDescription(), res.GetDone())
	}
}
