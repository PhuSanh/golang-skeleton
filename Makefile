include .env

GOCMD=go

dev-build:
	docker-compose -f docker-compose.local.yml up --build

prd-build:
	docker-compose build

dev-run:
	docker-compose -f docker-compose.local.yml up

prd-run:
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