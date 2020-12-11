# Go Gin App

This is a production ready go-gin scaffold project, which have below features out of box, helps you kick start a new project with minimum effort.

- multi environment config support
- full featured logging service
- mysql & redis support
- static file server example
- sample database models
- sample restful api
- user session management
- api response encapsulation
- timezone in consideration
- a dockerfile for deploy
- custom middleware example
- distributed lock implementation
- record soft delete feature

## Project Structure

```
go-gin-app
├── common              # common utils  
├── config              # config files
├── doc                 # documents
├── global              # global const / error definition
├── go.mod
├── go.sum
├── handler             # entrance for router
├── main.go             # init & start project
├── middleware          # middlewares
├── model               # db models & sql operation
├── public              # public static files
├── router              # router for api
├── run-dev             # scripts for run program
├── service             # main business logic implementation
└── view                # request / response definition
```

## Sample config file

```
env: dev
server:
  port: 8080
  timezoneLoc: Asia/Shanghai
  # gin mode: debug, release, test
  ginMode: debug
db:
  user: root
  pass: root
  host: 127.0.0.1
  port: 3306
  name: test
  maxConnect: 10
  maxIdle: 10
  showSql: true
redis:
  host: 127.0.0.1:6379
  pass: pass
  db: 0
log:
  path: /tmp/log
  # 0-PanicLevel 1-FatalLevel 2-ErrorLevel 3-WarnLevel 4-InfoLevel 5-DebugLevel 6-TraceLevel
  level: 5
```

## Quick Start

Install [Go](https://golang.org/dl/) and enable `GO111MODULE`

```bash
export GO111MODULE=on
```

Fetch and install dependencies listed in go.mod

```bash
go build ./...
```

Docker run mysql container 

```bash
sudo docker run -d --restart always -p 3306:3306 --name mysql5.7 -v /var/lib/mysql:/var/lib/mysql -e MYSQL_ROOT_HOST=% -e MYSQL_ROOT_PASSWORD=root mysql:5.7
```

Docker run redis container 

```bash
docker run -d --restart always --name redis -p 6379:6379 redis --requirepass "pass"
```

Create test database & table

```
mysql -uroot -p -h127.0.0.1 < doc/db/db.sql
```

To run your application locally:

```bash
go run main.go
```

To run specific environment by flag

```bash
go run main.go --env=sit
```

Or you can have standalone configuration file

```bash
go run main.go --config=/path/to/config.yml
```

## License

This sample application is licensed under the Apache License, Version 2.
