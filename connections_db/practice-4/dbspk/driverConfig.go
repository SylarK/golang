package dbspk

import "fmt"

const PostgresDriver = "postgres"

const User = "SPK"

const Host = "localhost"

const Port = "5432"

const Password = "-"

const DbName = "BRUTAMADEIRAS"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)