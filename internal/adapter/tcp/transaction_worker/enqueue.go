package transaction_worker

import (
	"time"

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

	req.Timestamp = time.Now()
	var balance int64 = account.Balance
	if req.Type == "c" {
		if req.Amount > account.Limit {
			req.Status = 0
			ok = false
		} else {
			req.Status = 1
			ok = true
		}
	} else {
		if (account.Balance - req.Amount) < -account.Limit {
			req.Status = 0
			ok = false
		} else {
			account.Balance -= req.Amount
			req.Status = 1
			ok = true
		}
		balance -= req.Amount
	}

	account.Push(&domain.Transaction{
		Type:        req.Type,
		Amount:      req.Amount,
		Description: req.Description,
		CreatedAt:   req.Timestamp,
	})

	w.ch <- req
	return account.Account.Limit, balance, ok, true
}
