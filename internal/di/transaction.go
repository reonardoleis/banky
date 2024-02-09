package di

import (
	"github.com/jmoiron/sqlx"

	"github.com/reonardoleis/banky/internal/adapter/http/transaction_service"
	"github.com/reonardoleis/banky/internal/adapter/postgres/transaction_repository"
	"github.com/reonardoleis/banky/internal/core/domain"
	"github.com/reonardoleis/banky/internal/core/usecase/transaction_usecases"
)

func Transaction(db *sqlx.DB) domain.TransactionService {
	repository := transaction_repository.New(db)
	usecases := transaction_usecases.New(repository)
	return transaction_service.New(usecases)
}
