package transaction_service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/reonardoleis/banky/internal/core/utils"
)

func (s service) GetStatement(ctx *fasthttp.RequestCtx) {
	path := strings.Split(string(ctx.Path()), "/")
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

	statement, err := s.usecases.GetStatement(uint(accountId))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "internal_server_error")
		return
	}

	utils.RespondWithJSON(ctx, statement.ToJSON())
}
