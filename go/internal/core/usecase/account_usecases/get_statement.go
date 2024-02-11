package account_usecases

import (
	"time"

	"github.com/reonardoleis/banky/internal/core/domain"
)

func (u usecase) GetStatement(
	userId uint,
	accounts map[uint]*domain.AccountInformation,
) (*domain.Statement, bool, error) {
	statement := new(domain.Statement)
	statement.Balance = new(domain.StatementBalance)
	statement.Transactions = make([]*domain.StatementTransaction, 0)

	account, ok := accounts[userId]
	if !ok {
		return nil, false, nil
	}

	statement.Balance.Total = account.Balance
	statement.Balance.Limit = account.Account.Limit
	statement.Balance.At = time.Now()

	for i := len(account.LastTransactions) - 1; i >= 0; i-- {
		tx := account.LastTransactions[i]
		statement.Transactions = append(
			statement.Transactions,
			&domain.StatementTransaction{
				Amount:      tx.Amount,
				Type:        tx.Type,
				Description: tx.Description,
				CreatedAt:   tx.CreatedAt,
			},
		)
	}

	return statement, true, nil
}
