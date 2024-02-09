package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/reonardoleis/banky/internal/adapter/http"
	"github.com/reonardoleis/banky/internal/adapter/postgres"
)

func main() {
	godotenv.Overload(".env")

	err := postgres.Connect()
	if err != nil {
		log.Fatal("error starting database", err)
	}

	if err := http.Run(); err != nil {
		log.Fatal("error running http server", err)
	}
}
