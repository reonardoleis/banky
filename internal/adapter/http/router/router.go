package router

import (
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/reonardoleis/banky/internal/adapter/postgres"
	"github.com/reonardoleis/banky/internal/di"
)

func HandleRoutes(ctx *fasthttp.RequestCtx) {
	transactionService := di.Transaction(postgres.DB())

	method := strings.ToLower(string(ctx.Method()))
	if method == "post" {
		transactionService.Create(ctx)
	} else if method == "get" {
		transactionService.GetStatement(ctx)
	} else {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "internal_server_error")
	}
}
