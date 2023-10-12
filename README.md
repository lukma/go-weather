# Weather API

This is **Golang** project for Weather WEB.

## Building the project

Setup locally

* Resolve project dependencies by run command `go mod tidy`.
* Put file `config/config.yml` contains require configuration for application inside config directory. Here example of config.yml
```
server_address: :8080
context_timeout: 2
db_config:
  host: AppDB
  port: 3306
  user: your_user
  pass: your_pass
  name: altechomegadb
```
* Import db schemes into mysql database.

To run application

* If you just want to running in terminal session just run command `go run cmd/apigw/main.go`
* If you just want to running in docker container just run command `docker-compose up`

To run tests

`go test -v ./...`

To generate code coverage reports

`go test -covermode=count -coverpkg=./... -coverprofile coverage.out -v ./...`
