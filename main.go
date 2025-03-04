package main

import (
	"log"
	"os"

	"app.myriadflow.com/db"
	"app.myriadflow.com/server"
	"github.com/joho/godotenv"
)

func init() {
	if len(os.Getenv("HOST")) == 0 {
		godotenv.Load()
	}
	if _, err := db.Connect(); err != nil {
		log.Fatalln(err)
	}
	db.Init()
}

func main() {
	server.Router()
}
