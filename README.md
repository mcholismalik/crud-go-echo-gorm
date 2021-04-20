# crud-go-echo-gorm
This project is boilerplate to build API in Go Echo Framework
- PORT : 3000
- PATH : /api/boilerplate

## Installation

``` bash
# clone the repo
$ git clone 

# go into app's directory
$ cd my-project
```

## Build & Run

``` bash
# If u wanna run in development 
$ ENV=DEV go run main.go

# If u wanna build 
$ go build

# If u wanna run in production
$ ENV=DEV ./repo-name

#If u wanna run in docker
$ docker-compose up 
```

## Generate swagger documentation

install go swagger

``` bash
go get github.com/swaggo/swag

cd $GOPATH/src/github.com/swaggo/swag

go install cmd/swag

```

to generate API documentation

``` bash

swag init

```

to see the results, run app and access {{base_url}}/swagger/index.html

## Feature 
This boilerplate have default feature
- CRUD 
- Middleware 
- Auth
- Pagination
- Response
- Env
- Redis
- Elasticsearch

# Author
mcholismalik.official@gmail.com