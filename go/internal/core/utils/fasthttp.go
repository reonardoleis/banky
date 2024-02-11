package utils

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func RespondWithJSON(ctx *fasthttp.RequestCtx, s []byte) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	fmt.Fprintf(ctx, string(s))
}
