package di

import (
	"github.com/jmoiron/sqlx"
	http "github.com/reonardoleis/banky/internal/adapter/http/account_service"
	"github.com/reonardoleis/banky/internal/adapter/postgres/account_repository"
	tcp "github.com/reonardoleis/banky/internal/adapter/tcp/account_service"
	"github.com/reonardoleis/banky/internal/adapter/tcp/account_worker"
	"github.com/reonardoleis/banky/internal/core/domain"
	"github.com/reonardoleis/banky/internal/core/usecase/account_usecases"
)

func AccountApi(db *sqlx.DB) domain.AccountApiService {
	repository := account_repository.New(db)
	usecases := account_usecases.New(repository)
	return http.New(usecases)
}

func AccountManager(db *sqlx.DB) domain.AccountManagerService {
	repository := account_repository.New(db)
	usecases := account_usecases.New(repository)
	worker := account_worker.New(usecases)
	return tcp.New(worker, usecases)
}
