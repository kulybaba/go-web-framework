package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB = func() (db *gorm.DB) {
	LoadEnv()

	db, err := gorm.Open(GetEnv("DATABASE"), GetEnv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if err := db.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	db.SingularTable(true)

	return
}()

var RedisClient = func() *redis.Client {
	db, err := strconv.Atoi(GetEnv("REDIS_DB"))
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     GetEnv("REDIS_URL"),
		Password: GetEnv("REDIS_PASSWORD"),
		DB:       db,
	})

	_, err = client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	return client
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
