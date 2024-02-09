package transaction_service

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/reonardoleis/banky/internal/core/dto"
	"github.com/reonardoleis/banky/internal/core/utils"
)

func (s service) Create(ctx *fasthttp.RequestCtx) {
	req, err := dto.JsonToCreateTransactionRequest(ctx.PostBody())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, "bad_request")
		return
	}

	transaction, err := s.usecases.Create(req)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprint(ctx, "internal_server_error")
	}

	utils.RespondWithJSON(ctx, transaction.ToJSON())
}
