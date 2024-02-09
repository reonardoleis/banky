package main

import (
	"log"

	"github.com/valyala/fasthttp"

	"github.com/reonardoleis/banky/adapter/http/router"
)

func main() {
	if err := fasthttp.ListenAndServe(":8080", router.HandleRoutes); err != nil {
		log.Fatal("error starting http server", err)
	}
}
