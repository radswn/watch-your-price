#!/bin/bash

usage() {
    echo "usage: db-migrate.sh COMMAND"
    echo ""
    echo "  COMMANDS:"
    echo "  create <name> - create new migration script with given <name>"
    echo "  update - update the database to the most recent version"
    echo "  drop - drop the whole database schema"
    echo "  version - print the current version of database"
    echo "  goto <ver> - migrate the database to version <ver>"
    echo "  up <x> - migrate the database up <x> versions"
    echo "  down <x> - migrate the database down <x> versions"
}

create_migration () {
    docker run -it --rm -v "$PWD"/migrations:/migrations \
    migrate/migrate create -ext sql -dir /migrations -seq "$1"
}

goto_version () {
    docker run -it --rm -v "$PWD"/migrations:/migrations \
    --network host migrate/migrate -path=/migrations/ \
    -database 'mysql://sa:!QAZxsw2@tcp(localhost:3306)/mydb' goto "$1"
}

up_version () {
    docker run -it --rm -v "$PWD"/migrations:/migrations \
    --network host migrate/migrate -path=/migrations/ \
    -database 'mysql://sa:!QAZxsw2@tcp(localhost:3306)/mydb' up "$1"
}

down_version () {
    docker run -it --rm -v "$PWD"/migrations:/migrations \
    --network host migrate/migrate -path=/migrations/ \
    -database 'mysql://sa:!QAZxsw2@tcp(localhost:3306)/mydb' down "$1"
}

print_version () {
    docker run -it --rm -v "$PWD"/migrations:/migrations \
    --network host migrate/migrate -path=/migrations/ \
    -database 'mysql://sa:!QAZxsw2@tcp(localhost:3306)/mydb' version
}

drop_db () {
    docker run -it --rm -v "$PWD"/migrations:/migrations \
    --network host migrate/migrate -path=/migrations/ \
    -database 'mysql://sa:!QAZxsw2@tcp(localhost:3306)/mydb' drop
}

update_db () {
    docker run -it --rm -v "$PWD"/migrations:/migrations \
    --network host migrate/migrate -path=/migrations/ \
    -database 'mysql://sa:!QAZxsw2@tcp(localhost:3306)/mydb' up
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
    up)
        if [ -n "$2" ]; then
          up_version "$2"
        else
          usage
        fi
        ;;
    down)
        if [ -n "$2" ]; then
          down_version "$2"
        else
          usage
        fi
        ;;
    update)
      update_db
      ;;
    drop)
      drop_db
      ;;
    version)
         print_version
        ;;
    *)
        usage
        ;;
esac
