#!/bin/bash

usage() {
    echo "db-migrate.sh <COMMAND>"
    echo ""
    echo "COMMANDS:"
    echo "create <name> - create new migration script with given name"
}

create_migration () {
    docker run -it --rm -v "$PWD"/migrations:/migrations migrate/migrate create -ext sql -dir /migrations -seq "$1"
}

case "$1" in
    create)
        if [ -n "$2" ]; then
          create_migration "$2"
        else
          usage
        fi
        ;;
    *)
        usage
        ;;
esac
