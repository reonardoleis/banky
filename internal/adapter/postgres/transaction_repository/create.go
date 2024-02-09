package transaction_repository

import (
	"github.com/reonardoleis/banky/internal/core/domain"
	"github.com/reonardoleis/banky/internal/core/dto"
)

func (r repository) Create(req *dto.CreateTransactionRequest) (*domain.Transaction, error) {
	transaction := new(domain.Transaction)

	err := r.db.QueryRow(
		`INSERT INTO transactions (amount, type, description, created_at)
     VALUES ($1, $2, $3, $4) RETURNING *`,
		req.Amount, req.Type, req.Description, "NOW()",
	).Scan(
		transaction.ID,
		transaction.Type,
		transaction.Amount,
		transaction.Description,
		transaction.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
