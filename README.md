# Exemplo Clean Architecture com Go

Este projeto demonstra o uso da Clean Architecture em Go, implementando uma API para gerenciamento de empresas (companies). 


## Configuração do projeto

Para inicializar o módulo Go no seu projeto, execute o seguinte comando:

`go mod init nome-do-seu-projeto`

Em seguida execute os seguintes comandos para instalar as dependências necessárias:

`go get github.com/gorilla/mux`
`go get github.com/denisenkom/go-mssqldb`
`go get github.com/joho/godotenv`
`go get -u github.com/swaggo/swag/cmd/swag`
`go get -u github.com/swaggo/http-swagger`
`go get -u github.com/swaggo/gin-swagger`

Após configurar a aplicação, gere a documentação Swagger com os comandos:

`swag init`

Caso o arquivo principal da aplicação esteja em uma subpasta (como no caso deste projeto):
`swag init -g cmd/server/main.go`

Para compilar o projeto, use o comando:

`go build ./...`

Agora para rodar o servidor com o seguinte comando:
`go run cmd/server/main.go`

Para abrir o swagger basta acessar o seguinte endereço no seu navegador: [Swagger](http://localhost:8080/swagger/index.html)
