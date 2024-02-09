package transaction_service

import "github.com/reonardoleis/banky/internal/core/domain"

type service struct {
	worker   domain.TransactionWorker
	usecases domain.TransactionUseCases
}

func New(
	worker domain.TransactionWorker,
	usecases domain.TransactionUseCases,
) domain.TransactionManagerService {
	return &service{
		worker:   worker,
		usecases: usecases,
	}
}

func (s service) Worker() domain.TransactionWorker {
	return s.worker
}
