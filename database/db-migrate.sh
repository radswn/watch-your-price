#!/bin/bash

usage() {
    echo "db-migrate.sh <COMMAND>"
    echo ""
    echo "COMMANDS:"
    echo "create <name> - create new migration script with given <name>"
    echo "goto <ver> - migrate the database to version <ver>"
}

create_migration () {
    docker run -it --rm -v "$PWD"/migrations:/migrations migrate/migrate create -ext sql -dir /migrations -seq "$1"
}

goto_version () {
    docker run -it --rm -v "$PWD"/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database 'mysql://sa:!QAZxsw2@tcp(localhost:3306)/mydb' goto "$1"
}

case "$1" in
    create)
        if [ -n "$2" ]; then
          create_migration "$2"
        else
          usage
        fi
        ;;
    goto)
        if [ -n "$2" ]; then
          goto_version "$2"
        else
          usage
        fi
        ;;

    *)
        usage
        ;;
esac
