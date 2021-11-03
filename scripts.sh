#!/bin/bash

export POSTGRESQL_URL='postgres://josephnguyen:1qaz2wsx@localhost:5432/todo?sslmode=disable'

# migration
# up    - migrate up
# down  - migrate down
# clean - fix and force database version
# gen   - generate code from queries folder

if [ -n "$1" ]; then
    case "$1" in
    up) migrate -database ${POSTGRESQL_URL} -path db/schema up;;
    down) migrate -database ${POSTGRESQL_URL} -path db/schema down;;
    clean) migrate -path db/schema -database ${POSTGRESQL_URL} force 1;;
    gen) sqlc generate
    *) echo "no valid options found";;
    esac
fi