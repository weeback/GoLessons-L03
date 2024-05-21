# Lession 03

Đặt tên sau

## Server API (RESTful API)

- A RESTful API (Representational State Transfer) is an architectural style for an application program interface (API) that uses HTTP requests to access and use data.
It is based on representational state transfer (REST) technology, an architectural style and approach to communications often used in web services development.
- A RESTful API breaks down a transaction to create a series of small modules. Each module addresses a particular underlying part of the transaction.
This modularity provides developers with a lot of flexibility.
- A RESTful API explicitly takes advantage of HTTP methodologies defined by the RFC 2616 protocol. 
They use GET to retrieve a resource; PUT to change the state of or update a resource, which can be an object, file or block; POST to create that resource; and DELETE to remove it.

* Setup Project
```shell
# init go module
go mod init myapp

# create static_variable.go file
echo "package myapp

// Version is the version of the application
var Version string
" >> static_variable.go

# create new main.go file
mkdir -p cmd/v1; echo "package main
  
import (
  \"fmt\"
    \"myapp\"
)

var (
  bindAddr = "0.0.0.0:8080"
)

func main() {
    fmt.Printf(\"Server API\n------------------------------------\n\"+
    \"\tVersion: %s\n------------------------------------\n\", myapp.Version)
}
" >> cmd/v1/main.go
```
  
* Build app
```shell
go mod tidy;\
  go build -ldflags "-s -w -extldflags '-static' -X myapp.Version=beta-1.0.0" \
  -o bin/myapp \
  -trimpath cmd/v1/*.go
```

* Run build
```shell
./bin/myapp
```
