package main

import (
	"os"

	"app.myriadflow.com/db"
	"app.myriadflow.com/server"
	"github.com/joho/godotenv"
)

func init() {
	if len(os.Getenv("HOST")) == 0 {
		godotenv.Load()
	}
	dbS, err := db.Connect()
	if err != nil {
		panic(err)

	}
	db.Init(dbS)
}

func main() {
	server.Router()
}
