package account_service

import "github.com/reonardoleis/banky/internal/core/domain"

type service struct {
	worker   domain.AccountWorker
	usecases domain.AccountUseCases
}

func New(
	worker domain.AccountWorker,
	usecases domain.AccountUseCases,
) domain.AccountManagerService {
	return &service{
		worker:   worker,
		usecases: usecases,
	}
}

func (s service) Worker() domain.AccountWorker {
	return s.worker
}
