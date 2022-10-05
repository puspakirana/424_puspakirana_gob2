# Summary Sesi 08

## Swagger
** Step Swaggo: **

> mkdir swaggo
> cd swaggo
> go mod init swaggo

** Library: **
> go get -u github.com/swaggo/swag/cmd/swag
> go get -u github.com/swaggo/http-swagger
> go get -u github.com/alecthomas/template
> go get -u github.com/gorilla/mux

** Generate swaggo/docs: **
> swag init -g main.go

** Global: **
> go install github.com/swaggo/swag/cmd/swag

** MacOS: Add Path **
export PATH=$(go env GOPATH)/bin:$PATH
