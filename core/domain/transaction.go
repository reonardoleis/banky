package domain

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/reonardoleis/banky/core/dto"
)

type Transaction struct {
	ID          uint
	Type        string
	Amount      int64
	Description string
	CreatedAt   time.Time
}

func (t Transaction) ToJSON() []byte {
	return []byte(fmt.Sprintf(
		`{"id":%d,"type":"%s","amount":%d,"description":"%s","created_at":"%s"}`,
		t.ID, t.Type, t.Amount, t.Description, t.CreatedAt.String(),
	))
}

type TransactionService interface {
	Create(ctx *fasthttp.RequestCtx)
}

type TransactionUseCases interface {
	Create(req *dto.CreateTransactionRequest) (*Transaction, error)
}

type TransactionRepository interface {
	Create(req *dto.CreateTransactionRequest) (*Transaction, error)
}
