#!/usr/bin/env bash

FILE=$1

echo -e "-- +migrate Up\n\n\n-- +migrate Down" > databases/migrations/$(date +%s)-${FILE}.sql