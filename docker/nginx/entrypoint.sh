#!usr/bin/env bash

envsubst < /etc/nginx/conf.d/conf.template > /etc/nginx/conf.d/api.conf

nginx -g 'daemon off;'