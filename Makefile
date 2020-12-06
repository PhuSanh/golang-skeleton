include .env

GOCMD=go

dev-build:
	docker network create -d bridge ${APP_NAME}_network || true
	docker-compose -f docker-compose-dev.yml up --build

prd-build:
	docker network create -d bridge ${APP_NAME}_network || true
	docker-compose build

dev-run:
	docker network create -d bridge ${APP_NAME}_network || true
	docker-compose -f docker-compose-dev.yml up

prd-run:
	docker network create -d bridge ${APP_NAME}_network || true
	docker-compose up -d

stop:
	docker-compose stop

dbcreate:
	docker exec -it ${APP_NAME}_api bash ./tools/scripts/dbcreate.sh $(name)

dbmigrate:
	docker exec -it ${APP_NAME}_api bash ./tools/scripts/dbmigrate.sh $(env) $(ops) $(limit)

gen-doc:
	docker exec -it ${APP_NAME}_api swag init -g cmd/server.go

gen:
	docker exec -it ${APP_NAME}_api bash ./tools/scripts/genmock.sh

unit:
	docker exec -it ${APP_NAME}_api ${GOCMD} test ./...