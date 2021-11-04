package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	pb "github.com/todo_service/generated/todopb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could  not connect : %v", err)
	}
	defer cc.Close()

	c := pb.NewTodoServiceClient(cc)

	testCreateTodoList(c)
}

func testCreateTodoList(c pb.TodoServiceClient) {
	fmt.Println("Testing CreateTodoList rpc")
	req := &pb.CreateTodoListRequest{
		TodoList: &pb.TodoList{
			TodoListName: "List",
		},
	}
	res, err := c.CreateTodoList(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling CreateTodoList RPC: %v", err)
	}

	log.Println(res.GetResponse())
}
