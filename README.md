## Quick start
1. Search `golang-skeleton` in project and rename it to your project name
2. RUN `cp .env.example .env`
3. RUN `make dev-build` to download package and build docker  
    - Note: You can encounter mysql connection error in the first time. Please stop process and RUN `make dev-run`
4. Add `127.0.0.1 local.api.vn` to file /etc/hosts
5. Access `local.api.vn` for API
6. Access `local.api.vn/swagger/index.html` for swagger documentation


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