package http

import (
	"os"

	"github.com/valyala/fasthttp"

	"github.com/reonardoleis/banky/internal/adapter/http/router"
)

func Run() error {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}

	return fasthttp.ListenAndServe(port, router.HandleRoutes)
}
