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
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}
	db.Init()
}

func main() {
	server.Router()
}
