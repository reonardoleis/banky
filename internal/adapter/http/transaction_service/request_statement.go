package transaction_service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/reonardoleis/banky/internal/core/utils"
)

func (s service) RequestStatement(ctx *fasthttp.RequestCtx) {
	path := strings.Split(string(ctx.Path()), "/")
	path = path[1:]
	if len(path) != 3 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, "invalid_path_params")
		return
	}

	accountId, err := strconv.Atoi(path[1])
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, "invalid_account_id")
		return
	}

	statement, err := s.usecases.RequestStatement(uint(accountId))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "internal_server_error")
		return
	}

	if statement == nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		fmt.Fprintf(ctx, "account_not_found")
		return
	}

	utils.RespondWithJSON(ctx, statement.ToJSON())
}
