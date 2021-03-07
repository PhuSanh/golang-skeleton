#!/bin/sh

### Load .env
if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi

### Add nginx conf
#### Why put env variable directly in command: https://github.com/docker-library/docs/issues/496
envsubst '$$API_URL' < docker/nginx/template/api.template > docker/nginx/conf.d/api.conf
if [ ${APP_ENV} != 'local' ]; then
  envsubst '$$PMA_URL' < docker/nginx/template/pma.template > docker/nginx/conf.d/pma.conf
fi

### Add nginx basic auth
envsubst < docker/nginx/template/htpasswd.template > docker/nginx/auth/.htpasswd

### Add createdb in mysql
envsubst < docker/mysql/docker-entrypoint-initdb.d/createdb.template > docker/mysql/docker-entrypoint-initdb.d/createdb.sql

### Create docker network
docker network create -d bridge ${APP_NAME}_network || true
