package configs

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB = func() (db *sql.DB) {
	LoadEnv()

	db, err := sql.Open(GetEnv("DATABASE"), GetEnv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return
}()

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func GetEnv(name string) (env string) {
	env, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("Not set variable %s in .env", name)
	}
	return
}
