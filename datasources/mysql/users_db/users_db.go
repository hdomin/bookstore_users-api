package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MYSQL_USERNAME = "MYSQL_USERNAME"
	MYSQL_PASSWORD = "MYSQL_PASSWORD"
	MYSQL_HOST     = "MYSQL_HOST"
	MYSQL_SCHEMA   = "MYSQL_SCHEMA"
)

var (
	Client *sql.DB

	username = os.Getenv(MYSQL_USERNAME)
	password = os.Getenv(MYSQL_PASSWORD)
	host     = os.Getenv(MYSQL_HOST)
	schema   = os.Getenv(MYSQL_SCHEMA)
)

func init() {
	var err error

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		username, password, host, schema,
	)

	temp, err := sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	Client = temp

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("Databases successfully configurated")

}
