package di

import (
	"github.com/jmoiron/sqlx"

	http_transaction_service "github.com/reonardoleis/banky/internal/adapter/http/transaction_service"
	"github.com/reonardoleis/banky/internal/adapter/postgres/transaction_repository"
	"github.com/reonardoleis/banky/internal/adapter/tcp/transaction_service"
	"github.com/reonardoleis/banky/internal/adapter/tcp/transaction_worker"
	"github.com/reonardoleis/banky/internal/core/domain"
	"github.com/reonardoleis/banky/internal/core/usecase/transaction_usecases"
)

func TransactionApi(db *sqlx.DB) domain.TransactionApiService {
	repository := transaction_repository.New(db)
	usecases := transaction_usecases.New(repository)
	return http_transaction_service.New(usecases)
}

func TransactionManager(db *sqlx.DB) domain.TransactionManagerService {
	repository := transaction_repository.New(db)
	usecases := transaction_usecases.New(repository)
	worker := transaction_worker.New(usecases)
	return transaction_service.New(worker, usecases)
}
