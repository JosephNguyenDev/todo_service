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
	testGetTodosFromList(c)
	testGetTodoById(c)
	testUpdateTodoNameById(c)
	testUpdateTodoListNameById(c)
	testUpdateTodoContentById(c)
	testUpdateTodoStatusById(c)
	testDeleteTodoById(c)
	testDeleteTodoListById(c)
	testGetTodoLists(c)
	testGetAllTodos(c)
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

func testDeleteTodoById(c pb.TodoServiceClient) {
	fmt.Println("testing DeleteTodoById rpc")
	req := &pb.DeleteTodoByIdRequest{
		Id: 1,
	}
	_, err := c.DeleteTodoById(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling DeleteTodoById RPC: %v", err)
	}
}

func testDeleteTodoListById(c pb.TodoServiceClient) {
	fmt.Println("testing DeleteTodoListById rpc")
	req := &pb.DeleteTodoListByIdRequest{
		Id: 1,
	}
	_, err := c.DeleteTodoListById(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling DeleteTodoListById RPC: %v", err)
	}
}

func testGetTodoById(c pb.TodoServiceClient) {
	fmt.Println("testing GetTodoById rpc")
	req := &pb.GetTodoByIdRequest{
		Id: 1,
	}
	res, err := c.GetTodoById(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GetTodosFromList RPC: %v", err)
	}
	log.Println(res.GetResponse())
}

func testGetAllTodos(c pb.TodoServiceClient) {
	fmt.Println("Testing GetAllTodos rpc")
	res, err := c.GetAllTodos(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling GetAllTodos RPC: %v", err)
	}

	log.Println(res.GetResponse())
}

func testUpdateTodoNameById(c pb.TodoServiceClient) {
	fmt.Println("testing UpdateTodoNameById rpc")
	req := &pb.UpdateTodoNameByIdRequest{
		Id:   1,
		Name: "List Part 2, Electric boogaloo",
	}
	res, err := c.UpdateTodoNameById(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling UpdateTodoNameById RPC: %v", err)
	}
	log.Println(res.GetResponse())
}

func testUpdateTodoContentById(c pb.TodoServiceClient) {
	fmt.Println("testing UpdateTodoContentById rpc")
	req := &pb.UpdateTodoContentByIdRequest{
		Id:      1,
		Content: "Some things to do, but more",
	}
	res, err := c.UpdateTodoContentById(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling UpdateTodoContetnById RPC: %v", err)
	}
	log.Println(res.GetResponse())
}

func testUpdateTodoStatusById(c pb.TodoServiceClient) {
	fmt.Println("testing UpdateTodoStatusById rpc")
	req := &pb.UpdateTodoStatusByIdRequest{
		Id:   1,
		Done: true,
	}
	res, err := c.UpdateTodoStatusById(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling UpdateTodoContentById RPC: %v", err)
	}
	log.Println(res.GetResponse())
}

func testUpdateTodoListNameById(c pb.TodoServiceClient) {
	fmt.Println("testing UpdateTodoListNameById rpc")
	req := &pb.UpdateTodoListNameByIdRequest{
		Id:   1,
		Name: "List Part 2, Electric boogaloo",
	}
	res, err := c.UpdateTodoListNameById(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling UpdateTodoNameById RPC: %v", err)
	}
	log.Println(res.GetResponse())
}
