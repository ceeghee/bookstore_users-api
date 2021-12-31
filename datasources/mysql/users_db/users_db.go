package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client   *sql.DB
	Test     string
	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	Client = connect()
}

func connect() (db *sql.DB) {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&tls=false",
		username, password, host, schema,
	)
	var err error
	dbClient, err := sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err := dbClient.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database succesfuly configured")
	return dbClient
}
