Author: [Nguyen Minh Tri](https://github.com/sirt102)

# Description

This is server site for Trinity system.

_Note:_

- **In this version, not have middleware yet. So all verify steps will be skiped and temporary use user_id in request body to detect request owner.**
- **In next version, the middleware layer and completely authentication - authorization module will be added. The verify steps will be change to use Token detacthed in Header Authorization of the request.**

# Document

- Design System: https://docs.google.com/document/d/1lpocPUvnas0_9uV2_Vg1LYGIuPA5rvpEHozHcb300_E/edit?usp=sharing
- Database Schema: [db schema](https://drive.google.com/file/d/1oPXOGL3uW0dk9ho0q6BdJO2uziuei2Hl/view?usp=drive_link)

# Infrastructure

- Golang - Gin
- PostgresQL
- Docker

# Folder structure

```
├── Dockerfile
├── README.md
├── cmd
│   └── server
│       └── main.go
├── config
│   └── config.go
├── docker-compose.yml
├── global
│   └── global.go
├── go.mod
├── go.sum
├── internal
│   ├── entities
│   │   ├── campaign.go
│   │   ├── common.go
│   │   ├── requests
│   │   │   ├── campaignrequest.go
│   │   │   ├── subscriptionrequest.go
│   │   │   └── userrequest.go
│   │   ├── role.go
│   │   ├── subscription.go
│   │   ├── transaction.go
│   │   ├── user.go
│   │   ├── usersubscription.go
│   │   ├── uservoucher.go
│   │   └── voucher.go
│   ├── handlers
│   │   ├── campaign.go
│   │   ├── subscription.go
│   │   ├── user.go
│   │   └── voucher.go
│   ├── initialize
│   │   ├── config.go
│   │   ├── initialize.go
│   │   ├── logger.go
│   │   ├── postgresql.go
│   │   ├── redis.go
│   │   └── router.go
│   ├── presenters
│   │   └── orderPresenter.go
│   ├── repositories
│   │   ├── campaign.go
│   │   ├── subscription.go
│   │   ├── transaction.go
│   │   ├── user.go
│   │   ├── usersubscription.go
│   │   ├── uservoucher.go
│   │   └── voucher.go
│   ├── routers
│   │   ├── admin
│   │   │   ├── campaign.go
│   │   │   ├── index.go
│   │   │   └── voucher.go
│   │   ├── public
│   │   │   ├── index.go
│   │   │   ├── subscription.go
│   │   │   └── user.go
│   │   └── router.go
│   ├── services
│   │   ├── campaign.go
│   │   ├── subscription.go
│   │   ├── user.go
│   │   └── voucher.go
│   ├── utils
│   │   ├── common
│   │   │   └── common.go
│   │   └── random
│   │       └── random.go
│   └── wire
│       ├── campaign.go
│       ├── subscription.go
│       ├── user.go
│       ├── voucher.go
│       └── wire_gen.go
├── logs
│   ├── README.md
│   └── dev.xxx.log
└── pkg
    └── logger
        └── logger.go
```

# How to run

**Requisite**

- Have Golang (use for local run) and Docker is running in local machine.
- Change Directory (cd command) to the root of project folder.
- Prepare .env file
- Prepare go swagger: `go get -u github.com/swaggo/swag `

  `cp .env.example .env`

  Fill the value to .env file to setup

- Could use the sample [.env](https://gist.github.com/sirt102/2e8988639cb3f4ab50ea005ffbf6ca35)
  If run in local: make sure set POSTGRESQL_HOST and REDIS_HOST values to local value: 127.0.0.1 or localhost
  If run in Docker: make sure set POSTGRESQL_HOST and REDIS_HOST values to host which defined in Docker Compose file

## Run Golang in local

`docker-compose down -v && docker-compose --build redis db -d`

`go mod download && go run ./cmd/server/main.go`

## Run Golang in Docker

`docker-compose down -v && docker-compose --build -d`

# Swagger - API Document

Run `swag init -g ./cmd/server/main.go -o ./cmd/swag/docs`

Swagger Endpoint: /swagger/index.html
