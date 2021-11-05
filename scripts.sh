#!/bin/bash

export POSTGRESQL_URL='postgres://josephnguyen:1qaz2wsx@localhost:5432/todo?sslmode=disable'

help=$'
up     - migrate up
down   - migrate down
clean  - fix and force database version
gen    - generate go code from queries folder
proto  - generate go code from proto
init   - initialize go module
help   - display available options
server - run testing server
test   - clear database and run testing client
evans  - run grpc testing cli
'

if [ -n "$1" ]; then
    case "$1" in
    up) migrate -database ${POSTGRESQL_URL} -path db/schema up;;
    down) migrate -database ${POSTGRESQL_URL} -path db/schema down;;
    clean) migrate -path db/schema -database ${POSTGRESQL_URL} force 1;;
    gen) sqlc generate;;
    proto) protoc --go_out=generated --go-grpc_out=generated proto/todo.proto;;
    init) 
        go mod init; 
        go mod tidy;;
    help) echo "$help";;
    server) go run testing/server/server.go;;
    test) 
        migrate -path db/schema -database ${POSTGRESQL_URL} force 1;
        migrate -database ${POSTGRESQL_URL} -path db/schema down <<< 'y';
        migrate -database ${POSTGRESQL_URL} -path db/schema up;
        go run testing/client/client.go;;
    evans) evans -p 50051 -r;;
    *) echo "no valid options found for $1
run './scripts.sh help' for more options";;
    esac
fi