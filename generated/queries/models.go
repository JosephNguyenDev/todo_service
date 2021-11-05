// Code generated by sqlc. DO NOT EDIT.

package query

import (
	"database/sql"
)

type Todo struct {
	ID           int32          `json:"id"`
	TodoName     sql.NullString `json:"todo_name"`
	TodoListName sql.NullString `json:"todo_list_name"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
	Content      string         `json:"content"`
	Done         bool           `json:"done"`
}

type TodoList struct {
	ID           int32          `json:"id"`
	TodoListName sql.NullString `json:"todo_list_name"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
}