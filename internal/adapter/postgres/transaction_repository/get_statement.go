package transaction_repository

import (
	"log"

	"github.com/reonardoleis/banky/internal/core/domain"
)

func (r repository) GetStatement(accountId uint) (*domain.Statement, error) {
	statement := new(domain.Statement)
	statement.Balance = new(domain.StatementBalance)
	statement.Transactions = make([]*domain.StatementTransaction, 0)

	err := r.db.QueryRow(
		`SELECT
     (accounts.initial_balance - SUM(transactions.amount)) as total,
      accounts.limit as limit,
      NOW() as at
      FROM accounts
      JOIN transactions ON (transactions.account_id = accounts.id)
      WHERE accounts.id = $1
      GROUP BY (accounts.id)`,
		accountId,
	).Scan(
		&statement.Balance.Total,
		&statement.Balance.Limit,
		&statement.Balance.At,
	)
	if err != nil {
		log.Println("transaction repository: error while getting statement", err)
		return nil, err
	}

	rows, err := r.db.Query(
		`SELECT amount, type, description, created_at
     FROM   transactions
     WHERE  account_id = $1
     AND    transactions.status = $2`,
		accountId, 1,
	)
	if err != nil {
		log.Println("transaction repository: error while getting statement", err)
		return nil, rows.Err()
	}

	p := 0
	defer rows.Close()
	for rows.Next() {
		statement.Transactions = append(statement.Transactions, &domain.StatementTransaction{})
		err = rows.Scan(
			&statement.Transactions[p].Amount,
			&statement.Transactions[p].Type,
			&statement.Transactions[p].Description,
			&statement.Transactions[p].CreatedAt,
		)
		if err != nil {
			log.Println("transaction repository: error whille parsing rows on get statement", err)
			return nil, err
		}
		p++
	}

	return statement, nil
}
