### Setup
1. Update file .env
``` 
cp .env.example .env
``` 
2. Setup before start
```
bash script/setup/before.sh
```
- Setup nginx conf
- Setup nginx basic auth for phpmyadmin
- Create init db file
3. Add host in /etc/hosts
```
127.0.0.1 {server_name} {pma_server_name}
```
4. Run docker compose
```
docker-compose -f docker-compose.prod.yml up -d
```
5. Setup after start
```
docker exec -it sev_phalcon bash script/setup/after.sh
```
- Install go library
- Run migration

Note: remember to enable go mod in Goland for development


## Start
##### Build container
```shell script
## Development
make dev-build

## Production
make prd-build
```

##### Run container
```shell script
## Development
make dev-run

## Production
make prd-run
```

## Migration
##### Create migration file
```shell script
## Outside container
make dbcreate name=[file name]

## Inside container
bash ./tools/scripts/dbcreate [file name]
```

##### Migrate
```shell script
## Outside container
make dbmigrate env=[development,production] ops=[status,up,down] limit=[#limit]

## Inside container
bash ./tools/scripts/dbmigrate.sh [development,production] [status,up,down] [#limit]
```

#### Add password for superadmin user
Because create user func need Auth, please add user Superadmin to database manually. You can run test `TestHashBCrypt` to get hashed password

Example: raw (superA@999), hashed ($2a$10$lYBqhxCNvtQUNiBneT.L..JNFe7ipLLtB.K0D8c8xi/wAWELaDqaK)


## Documentation
Access [host]:[port]/swagger/index.html
```shell script
## Outside container
make gen-doc

## Inside container
swag init -g cmd/server.go
```

## TODO
- [ ] Domain Driven Design pattern
- [x] Viper read config
- [x] Server Echo
- [x] GORM database
- [x] Redis
- [ ] Logging
    - [x] Init Uber zap
    - [ ] Log to file
- [x] JWT
- [x] Migration (https://github.com/rubenv/sql-migrate)
- [x] Queue
- [ ] Testing
    - [x] Mockery
    - [x] Unit test
    - [ ] Integration test
- [x] Swagger documentation
- [ ] Cronjob
- [ ] Send mail (sendgrid)
- [x] Docker
    - [x] Multi-stage prod
    - [x] Auto-reload dev
- [ ] Gitlab CI
- [ ] Linter / Goimports