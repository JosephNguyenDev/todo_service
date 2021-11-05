package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	querier "github.com/todo_service/generated/queries"
	pb "github.com/todo_service/generated/todopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

// global database variable
var DB *querier.Queries

type server struct {
	pb.UnimplementedTodoServiceServer
}

func (*server) CreateTodoList(ctx context.Context, req *pb.CreateTodoListRequest) (*pb.CreateTodoListResponse, error) {
	res, err := DB.CreateTodoList(
		ctx,
		sql.NullString{
			String: req.TodoList.GetTodoListName(),
			Valid:  true},
	)

	if err != nil {
		return nil, err
	}

	s, err := json.MarshalIndent(res, "", " ")
	return &pb.CreateTodoListResponse{
		Response: string(s),
	}, err
}

func (*server) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	res, err := DB.CreateTodo(
		ctx,
		querier.CreateTodoParams{
			TodoName: sql.NullString{
				String: req.Todo.GetTodoName(),
				Valid:  true,
			},
			TodoListName: sql.NullString{
				String: req.Todo.GetTodoListName(),
				Valid:  true,
			},
			Content: req.Todo.Content,
			Done:    req.Todo.Done,
		},
	)

	if err != nil {
		return nil, err
	}

	s, err := json.MarshalIndent(res, "", " ")
	return &pb.CreateTodoResponse{
		Response: string(s),
	}, err
}

func (*server) GetTodosFromList(ctx context.Context, req *pb.GetTodosFromListRequest) (*pb.GetTodosFromListResponse, error) {
	res, err := DB.GetTodosFromList(
		ctx,
		sql.NullString{
			String: req.GetList(),
			Valid:  true},
	)
	if err != nil {
		return nil, err
	}
	s, err := json.MarshalIndent(res, "", " ")
	return &pb.GetTodosFromListResponse{
		Response: string(s),
	}, err
}

func (*server) GetTodoLists(ctx context.Context, _ *emptypb.Empty) (*pb.GetTodoListsResponse, error) {
	res, err := DB.GetTodoLists(ctx)
	if err != nil {
		return nil, err
	}
	s, err := json.MarshalIndent(res, "", " ")
	return &pb.GetTodoListsResponse{
		Response: string(s),
	}, err
}

func main() {
	// set up listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	// Open database connection
	conn, err := sql.Open("postgres", "user=josephnguyen password=1qaz2wsx dbname=todo sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	db := querier.New(conn)
	DB = db
	fmt.Println("Established database connection")

	// start server
	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &server{})
	fmt.Println("Server Started")

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
