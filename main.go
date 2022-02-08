package main

import (
	"github.com/swkkd/crud/database"
	"github.com/swkkd/crud/routers"
	"log"
	"os"
)

var DSN string

func init() {
	DB_HOST := "host=" + readFromENV("DB_HOST", "localhost") + " "
	DB_PORT := "port=" + readFromENV("DB_PORT", "5432") + " "
	DB_USER := "user=" + readFromENV("DB_USER", "postgres") + " "
	DB_NAME := "dbname=" + readFromENV("DB_NAME", "mydb") + " "
	DB_SSLMODE := "sslmode=" + readFromENV("DB_SSLMODE", "disable") + " "
	DB_PASSWORD := "password=" + readFromENV("DB_PASSWORD", "root") + " "
	DSN = DB_HOST + DB_PORT + DB_USER + DB_NAME + DB_SSLMODE + DB_PASSWORD
}

func main() {
	database.Connection(DSN)
	routers.Setup()
	sqlDb, err := database.DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDb.Close()
}

func readFromENV(key, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}
