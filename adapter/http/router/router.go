package router

import (
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/reonardoleis/banky/adapter/postgres"
	"github.com/reonardoleis/banky/di"
)

func HandleRoutes(ctx *fasthttp.RequestCtx) {
	transactionService := di.Transaction(postgres.DB())

	method := strings.ToLower(string(ctx.Method()))
	if method == "post" {
		transactionService.Create(ctx)
	} else if method == "get" {
		fmt.Fprintf(ctx, "not_implemented")
	} else {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "internal_server_error")
	}
}
