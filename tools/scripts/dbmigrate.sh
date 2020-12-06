#!/usr/bin/env bash

ENV=$1
OPS=$2
LIMIT=${3:-1}

if [ "$OPS" != "status" ] && [ "$OPS" != "up" ] && [ "$OPS" != "down" ] || [ "$ENV" != "development" ] && [ "$ENV" != "production" ]
then
  echo "Usage: bash ./dbmigrate.sh [development,production] [status,up,down] [#limit]"
  exit 1
fi

# Source file .env so that dbconfig.yml can use env variable
# https://gist.github.com/mihow/9c7f559807069a03e302605691f85572
export $(grep -v '^#' .env | xargs)

if [ "$OPS" == "status" ]
then
  sql-migrate $OPS -config=databases/migrations/dbconfig.yml -env=$ENV
else
  sql-migrate $OPS -config=databases/migrations/dbconfig.yml -env=$ENV -limit=$LIMIT
fi