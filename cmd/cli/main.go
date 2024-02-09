package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/reonardoleis/banky/internal/adapter/http"
	"github.com/reonardoleis/banky/internal/adapter/postgres"
	"github.com/reonardoleis/banky/internal/adapter/tcp"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("missing command")
	}

	godotenv.Overload(".env")

	err := postgres.Connect()
	if err != nil {
		log.Fatalln("error starting database", err)
	}

	if args[0] == "tcp" {
		if err := tcp.Run(); err != nil {
			log.Fatalln("error while running tcp adapter", err)
		}
	} else if args[0] == "http" {
		if err := http.Run(); err != nil {
			log.Fatalln("error while running http adapter", err)
		}
	} else {
		panic("invalid command")
	}

}
