package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	DB_DRIVER   string = "mysql" // "mysql" or "postgres"
	DB_USER     string = "username"
	DB_PASSWORD string = "password"
	DB_HOST     string = "localhost"
	DB_PORT     string = "3306"
	DB_NAME     string = "db_name"
)

var DB = func() *sql.DB {
	db, err := sql.Open(DB_DRIVER, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME))
	if err != nil {
		log.Fatal(err)
	}
	return db
}()
