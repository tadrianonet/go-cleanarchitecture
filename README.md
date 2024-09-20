init modulo
`go mod init nome-do-seu-projeto`

install
`go get github.com/gorilla/mux`
`go get github.com/denisenkom/go-mssqldb`
`go get github.com/joho/godotenv`

`go get -u github.com/swaggo/swag/cmd/swag`
`go get -u github.com/swaggo/http-swagger`
`go get -u github.com/swaggo/gin-swagger`


build
`go build ./...`

swagger
`swag init`
`swag init -g cmd/server/main.go`

run
go run cmd/server/main.go
