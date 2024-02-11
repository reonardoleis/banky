package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/reonardoleis/banky/internal/adapter/postgres"
	"github.com/reonardoleis/banky/internal/adapter/tcp"
)

func main() {
	godotenv.Overload(".env")
	if err := postgres.Connect(); err != nil {
		log.Fatalln("error starting database", err)
	}

	if err := tcp.Run(); err != nil {
		log.Fatalln("error while running tcp adapter", err)
	}
}
