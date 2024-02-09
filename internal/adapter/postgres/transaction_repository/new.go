package transaction_repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/reonardoleis/banky/internal/core/domain"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) domain.TransactionRepository {
	return &repository{
		db: db,
	}
}
