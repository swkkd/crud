package main

import (
	"github.com/swkkd/crud/database"
	"github.com/swkkd/crud/routers"
	"log"
)

func main() {
	database.Connection()
	routers.Setup()
	sqlDb, err := database.DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDb.Close()

}
