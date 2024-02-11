package account_usecases

import (
	"log"

	"github.com/reonardoleis/banky/internal/core/dto"
)

func (u usecase) CreateTransaction(req []*dto.CreateTransactionRequest) error {
	err := u.r.CreateTransaction(req)
	if err != nil {
		log.Println("transaction usecase: error creating transaction", err)
		return err
	}

	return nil
}
