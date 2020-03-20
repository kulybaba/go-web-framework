package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB = func() (db *gorm.DB) {
	LoadEnv()

	db, err := gorm.Open(GetEnv("DATABASE"), GetEnv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return
}()

func LoadEnv() {
	if err := godotenv.Load(".env.local"); err != nil {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal(err)
		}
	}
}

func GetEnv(name string) (env string) {
	env, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("Not set variable %s in .env", name)
	}
	return
}
