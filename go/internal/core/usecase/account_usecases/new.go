package account_usecases

import "github.com/reonardoleis/banky/internal/core/domain"

type usecase struct {
	r domain.AccountRepository
}

func New(
	repository domain.AccountRepository,
) domain.AccountUseCases {
	return &usecase{repository}
}
