package account_repository

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/reonardoleis/banky/internal/core/dto"
)

func (r repository) CreateTransaction(req []*dto.CreateTransactionRequest) error {
	query := `INSERT INTO transactions (status, account_id, amount, "type", description, created_at) VALUES `

	for _, transaction := range req {
		placeholders := fmt.Sprintf("('%d', %d, %d, '%s', '%s', '%s')",
			transaction.Status,
			transaction.AccountId,
			transaction.Amount,
			transaction.Type,
			transaction.Description,
			transaction.Timestamp.Format(time.RFC3339))

		query += placeholders + ","
	}

	query = strings.TrimSuffix(query, ",")
	_, err := r.db.Exec(query)
	if err != nil {
		log.Println("transaction repository: error creating transactions", err)
		return err
	}

	return nil
}
