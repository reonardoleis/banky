package transaction_repository

import (
	"github.com/reonardoleis/banky/internal/core/domain"
)

func (r repository) GetStatement(accountId uint) (*domain.Statement, error) {
	statement := new(domain.Statement)

	err := r.db.QueryRow(
		`SELECT 
      (accounts.initial_balance - SUM(transactions.amount)) as total
      accounts.limit as limit
      NOW() as at
      WHERE accounts.id = $1
      AND   transactions.account_id = accounts.id`,
		accountId,
	).Scan(
		statement.Balance.Total,
		statement.Balance.Limit,
		statement.Balance.At,
	)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(
		`SELECT amount, type, description, created_at
     FROM   transactions
     WHERE  account_id = $1`,
		accountId,
	)
	if err != nil {
		return nil, rows.Err()
	}

	defer rows.Close()
	for rows.Next() {
		statement.Transactions = append(statement.Transactions, &domain.StatementTransaction{})
		err = rows.Scan(
			statement.Transactions[len(statement.Transactions)-1].Amount,
			statement.Transactions[len(statement.Transactions)-1].Type,
			statement.Transactions[len(statement.Transactions)-1].Description,
			statement.Transactions[len(statement.Transactions)-1].CreatedAt,
		)
	}

	return statement, nil
}
