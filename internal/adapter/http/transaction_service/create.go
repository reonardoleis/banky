package transaction_service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/reonardoleis/banky/internal/core/dto"
	"github.com/reonardoleis/banky/internal/core/utils"
)

func (s service) Create(ctx *fasthttp.RequestCtx) {
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

	req, err := dto.JsonToCreateTransactionRequest(ctx.PostBody())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, "bad_request")
		return
	}

	req.AccountId = uint(accountId)

	transaction, err := s.usecases.Create(req)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprint(ctx, "internal_server_error")
	}

	utils.RespondWithJSON(ctx, transaction.ToJSON())
}
