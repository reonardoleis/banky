package main

import (
	"log"

	"github.com/reonardoleis/banky/internal/adapter/http"
)

func main() {
	if err := http.Run(); err != nil {
		log.Fatal("error running http server", err)
	}
}
