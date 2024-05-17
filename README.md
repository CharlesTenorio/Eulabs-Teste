# Eulabs  PRODUCT API

Api Crud de produtos em Golang usando Echo-fremework

> Requisitos do projeto:

- Go Lang >= 1.18

As demais dependências estão no arquivo go.mod

- https://go.dev/dl/

> Build do Back-End Go:
```bash
# Baixando as dependências
$ go mod tidy

# Compilar servidor HTTP
$ go build -o main cmd/api/main.go

# Ou compilar para outra plataforma ex: windows
$ GOOS=windows GOARCH=amd64 go build -o main64.exe cmd/api/main.go


$ go build -ldflags "-s -w" .
# Ou
$ go build -ldflags "-s -w" cmd/api/main.go
# Ou
$ go build -ldflags "-s -w" -o main cmd/api/main.go
```
## Tem quer seta esssa vars de ambiente antes de rodar a aplicação No linux dentro da pasta cmd/api
export SRV_DB_DRIVE="mysql"
export SRV_DB_HOST="localhost"
export SRV_DB_PORT="3306"
export SRV_DB_USER="root"
export SRV_DB_PASS="supersenha"
export SRV_DB_NAME="eulabs_db_dev"
## Opções de execução
- SRV_PORT (Porta padrão 8080)
- SRV_MODE (developer, homologation ou production / padrão production)

> Exemplo de Uso:
```bash
$ ./main.exe
# Ou
$ SRV_PORT=8080 SRV_MODE=developer ./main.exe
# Ou
$ SRV_PORT=9090 SRV_MODE=production ./main.exe
```

> Acesse:
- http://localhost:8080/
### RODAR myqql no docker para dev
docker-compose up -d dbmysql
docker-compose down

#### Para rodar os endpoints utilize o Postman
https://www.postman.com/downloads/
apos baixa e instalar postaman abra e va na opção import escolha o arquivo
 ApiProduct.json na pasta docs/Insomnia so executar os endpoint 
