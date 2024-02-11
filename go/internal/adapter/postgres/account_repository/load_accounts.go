package account_repository

import (
	"log"

	"github.com/reonardoleis/banky/internal/core/domain"
)

func (r repository) LoadAccounts() (map[uint]*domain.AccountInformation, error) {
	accounts := make(map[uint]*domain.AccountInformation)

	rows, err := r.db.Query(
		`SELECT 
      accounts.id, 
      accounts.initial_balance, 
      accounts."limit", 
      (accounts.initial_balance 
      + COALESCE(SUM(CASE WHEN transactions.type = 'c' THEN transactions.amount ELSE 0 END), 0) 
      - COALESCE(SUM(CASE WHEN transactions.status = 1 AND transactions.type = 'd' THEN transactions.amount ELSE 0 END), 0)) AS balance
    FROM 
      accounts
    LEFT OUTER JOIN 
      transactions ON (transactions.account_id = accounts.id)
    GROUP BY 
      accounts.id, accounts.initial_balance, accounts."limit"`,
	)
	if err != nil {
		log.Println("account repository error while loading accounts", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		account := new(domain.AccountInformation)
		account.Account = new(domain.Account)
		account.LastTransactions = make([]*domain.Transaction, 0)
		account.Balance = 0
		err := rows.Scan(
			&account.Account.ID,
			&account.Account.InitialBalance,
			&account.Account.Limit,
			&account.Balance,
		)
		if err != nil {
			log.Println("account repository error while scanning loaded account", err)
			return nil, err
		}

		account.Limit = account.Account.Limit

		accounts[account.Account.ID] = account
	}

	return accounts, nil
}
