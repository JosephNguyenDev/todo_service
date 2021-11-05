package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	pb "github.com/todo_service/generated/todopb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer cc.Close()

	c := pb.NewTodoServiceClient(cc)
	testCreateTodoList(c)
	testCreateTodo(c)
	testGetTodoLists(c)
	testGetTodosFromList(c)
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

func testCreateTodo(c pb.TodoServiceClient) {
	fmt.Println("Testing CreateTodo rpc")
	req := &pb.CreateTodoRequest{
		Todo: &pb.Todo{
			TodoName:     "Task",
			TodoListName: "List",
			Content:      "This is a thing to do later",
			Done:         false,
		},
	}
	res, err := c.CreateTodo(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling CreateTodo RPC: %v", err)
	}

	log.Println(res.GetResponse())
}

func testGetTodoLists(c pb.TodoServiceClient) {
	fmt.Println("Testing GetTodoLists rpc")
	res, err := c.GetTodoLists(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling GetTodoLists RPC: %v", err)
	}

	log.Println(res.GetResponse())
}

func testGetTodosFromList(c pb.TodoServiceClient) {
	fmt.Println("Testing GetTodosFromList rpc")
	req := &pb.GetTodosFromListRequest{
		List: "List",
	}
	res, err := c.GetTodosFromList(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GetTodosFromList RPC: %v", err)
	}

	log.Println(res.GetResponse())
}
