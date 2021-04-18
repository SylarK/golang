// Arquivo sera usado como pacote e importado para o main

package dbconfig

import "fmt"

// Definindo tabela criada no banco para teste
type Article struct {
	ID    int
	Title string
	Body  []byte
}

const PostgresDriver = "postgres"

const User = "postgres"

const Host = "localhost"

const Port = "5432"

const Password = "masterkey"

const DbName = "Go"

var TableName = [...]string{"product", "users"}

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
