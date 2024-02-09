package transaction_service

import "github.com/reonardoleis/banky/core/domain"

type service struct {
	usecases domain.TransactionUseCases
}

func New(usecases domain.TransactionUseCases) domain.TransactionService {
	return &service{
		usecases: usecases,
	}
}
