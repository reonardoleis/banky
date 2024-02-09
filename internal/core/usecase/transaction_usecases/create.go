package transaction_usecases

import (
	"fmt"
	"time"

	"github.com/reonardoleis/banky/internal/core/domain"
	"github.com/reonardoleis/banky/internal/core/dto"
)

func (u usecase) Create(req *dto.CreateTransactionRequest) (*domain.Transaction, error) {
	//transaction, err := u.r.Create(req)
	//if err != nil {
	//	return nil, err
	//}

	fmt.Printf("transaction_usecases.usecase.Create(%+v)\n", req)

	return &domain.Transaction{
		ID:          1,
		Amount:      req.Amount,
		Type:        req.Type,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}, nil
}
