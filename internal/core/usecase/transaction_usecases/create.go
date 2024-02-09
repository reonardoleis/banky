package transaction_usecases

import (
	"log"

	"github.com/reonardoleis/banky/internal/core/dto"
)

func (u usecase) Create(req []*dto.CreateTransactionRequest) error {
	err := u.r.Create(req)
	if err != nil {
		log.Println("transaction usecase: error creating transaction", err)
		return err
	}

	return nil
}
