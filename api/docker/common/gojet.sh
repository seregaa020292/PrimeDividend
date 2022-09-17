#!/bin/bash

readonly DB_PASSWORD=$(cat $DB_PASSWORD_FILE)

dsn="user=$DB_USER password=$DB_PASSWORD host=$DB_HOST port=$DB_PORT dbname=$DB_NAME sslmode=disable"

jet -source=postgres -dsn="$dsn" "$@"
