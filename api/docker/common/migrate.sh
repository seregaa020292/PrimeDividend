#!/bin/bash

readonly DB_PASSWORD=$(cat $DB_PASSWORD_FILE)

export GOOSE_DBSTRING="host=$DB_HOST user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable"

goose -dir $GOOSE_DIR_MIGRATIONS "$@"
