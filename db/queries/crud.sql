-- name: CreateTodo :one
INSERT INTO todo (todo_name, todo_list_name, content, done)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: CreateTodoList :one
INSERT INTO todo_list (todo_list_name)
VALUES($1) RETURNING *;

-- name: GetTodosFromList :many
SELECT * FROM todo
WHERE todo_list_name = $1;

-- name: GetTodoLists :many
SELECT * FROM todo_list;

-- name: DeleteTodoById :exec
DELETE FROM todo
WHERE id = $1;

-- name: DeleteTodoListById :exec
DELETE FROM todo_list
WHERE id = $1;

-- name: GetTodoById :one
SELECT * FROM todo
WHERE id = $1;

-- name: GetAllTodos :many
SELECT * FROM todo;

-- name: UpdateTodoNameById :one
UPDATE todo
SET todo_name = $2
WHERE id = $1 RETURNING *;

-- name: UpdateTodoContentById :one
UPDATE todo
SET content = $2
WHERE id = $1 RETURNING *;

-- name: UpdateTodoStatusById :one
UPDATE todo
SET done = $2
WHERE id = $1 RETURNING *;

-- name: UpdateTodoListNameById :one
UPDATE todo_list
SET todo_list_name = $2
WHERE todo_list_name = $1 RETURNING *;