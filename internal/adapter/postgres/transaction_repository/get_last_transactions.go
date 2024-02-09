package transaction_repository

import "github.com/reonardoleis/banky/internal/core/domain"

func (r repository) GetLastTransactions(accountId uint) ([]*domain.Transaction, error) {
	rows, err := r.db.Query(
		`SELECT 
		id, 
		amount, 
		type, 
		description, 
		created_at
		FROM transactions
		WHERE account_id = $1 AND status = $2
		ORDER BY created_at DESC
		LIMIT 10`,
		accountId, 1,
	)
	if err != nil {
		return nil, err
	}

	transactions := make([]*domain.Transaction, 0)
	defer rows.Close()
	for rows.Next() {
		transaction := new(domain.Transaction)
		err := rows.Scan(
			&transaction.ID,
			&transaction.Amount,
			&transaction.Type,
			&transaction.Description,
			&transaction.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	for i, j := 0, len(transactions)-1; i < j; i, j = i+1, j-1 {
		transactions[i], transactions[j] = transactions[j], transactions[i]
	}

	return transactions, nil
}
