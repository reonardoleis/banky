package account_usecases

import (
	"time"

	"github.com/reonardoleis/banky/internal/core/domain"
	"github.com/reonardoleis/banky/internal/core/dto"
)

func (u usecase) HandleTransaction(
	accountInformation *domain.AccountInformation,
	req *dto.CreateTransactionRequest,
) (limit, balance int64, ok bool) {
	req.Timestamp = time.Now()

	if req.Type == "c" {
		accountInformation.Balance += req.Amount
		req.Status = 1
	} else {
		if (accountInformation.Balance - req.Amount) >= -accountInformation.Limit {
			accountInformation.Balance -= req.Amount
			req.Status = 1
		} else {
			req.Status = 0
		}
	}

	return accountInformation.Limit, accountInformation.Balance, req.Status == 1
}
