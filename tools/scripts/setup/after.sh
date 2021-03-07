#!/bin/sh

### Load .env
if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi

### Run migration
if [ ${APP_ENV} == 'production' ]; then
  bash ./tools/scripts/dbmigrate.sh production up 1000
else
  bash ./tools/scripts/dbmigrate.sh development up 1000
fi
