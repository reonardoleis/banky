package transaction_worker

import (
	"log"

	"github.com/reonardoleis/banky/internal/core/domain"
)

func (w worker) GetStatement(accountId uint) (*domain.Statement, bool, error) {
	statement, exists, err := w.usecase.GetStatement(accountId, accounts)
	if err != nil {
		log.Println("transaction worker error while getting statement", err)
		return nil, false, err
	}

	return statement, exists, nil
}
