package transaction_service

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/valyala/fasthttp"

	"github.com/reonardoleis/banky/internal/core/dto"
	"github.com/reonardoleis/banky/internal/core/utils"
)

func (s service) DispatchCreation(ctx *fasthttp.RequestCtx) {
	path := strings.Split(string(ctx.Path()), "/")
	path = path[1:]
	if len(path) != 3 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, "invalid_path_params")
		return
	}

	accountId, err := strconv.Atoi(path[1])
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, "invalid_account_id")
		return
	}

	req, err := dto.JsonToCreateTransactionRequest(ctx.PostBody())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, "bad_request")
		return
	}

	descriptionLength := utf8.RuneCountInString(req.Description)
	if descriptionLength == 0 || descriptionLength > 10 {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, "invalid_description_length")
		return
	}

	if req.Type != "c" && req.Type != "d" {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		fmt.Fprintf(ctx, "invalid_type")
		return
	}

	req.AccountId = uint(accountId)

	transaction, ok, err := s.usecases.DispatchCreation(req)
	if err != nil {
		log.Println("transaction service: error creating transaction", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprint(ctx, "internal_server_error")
		return
	}

	if transaction == nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		fmt.Fprintf(ctx, "account_not_found")
		return
	}

	if transaction.Balance < -transaction.Limit || !ok {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		if req.Type == "d" {
			transaction.Balance += req.Amount
		}
	}

	utils.RespondWithJSON(ctx, transaction.ToJSON())
}
