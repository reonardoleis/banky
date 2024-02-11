package account_worker

import (
	"github.com/reonardoleis/banky/internal/core/domain"
	"github.com/reonardoleis/banky/internal/core/dto"
)

func (w worker) Enqueue(req *dto.CreateTransactionRequest) (int64, int64, bool, bool) {
	lock, ok := w.locks[req.AccountId]
	if !ok {
		return 0, 0, false, false
	}

	lock.Lock()
	defer lock.Unlock()

	account := accounts[req.AccountId]

	w.usecase.HandleTransaction(account, req)

	account.Push(&domain.Transaction{
		Type:        req.Type,
		Amount:      req.Amount,
		Description: req.Description,
		CreatedAt:   req.Timestamp,
	})

	w.ch <- req
	return account.Limit, account.Balance, req.Status == 1, true
}
