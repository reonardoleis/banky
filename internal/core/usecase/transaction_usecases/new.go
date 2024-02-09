package transaction_usecases

import "github.com/reonardoleis/banky/internal/core/domain"

type usecase struct {
	r domain.TransactionRepository
}

func New(
	repository domain.TransactionRepository,
) domain.TransactionUseCases {
	return &usecase{repository}
}
