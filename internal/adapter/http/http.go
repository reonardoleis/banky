package http

import (
	"github.com/valyala/fasthttp"

	"github.com/reonardoleis/banky/internal/adapter/http/router"
)

func Run() error {
	return fasthttp.ListenAndServe(":8080", router.HandleRoutes)
}
