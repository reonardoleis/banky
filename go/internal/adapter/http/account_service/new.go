package account_service

import "github.com/reonardoleis/banky/internal/core/domain"

type service struct {
	usecases domain.AccountUseCases
}

func New(usecases domain.AccountUseCases) domain.AccountApiService {
	return &service{
		usecases: usecases,
	}
}
