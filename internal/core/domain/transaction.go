package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/reonardoleis/banky/internal/core/dto"
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

type StatementBalance struct {
	Total int64
	Limit int64
	At    time.Time
}

type StatementTransaction struct {
	Amount      int64
	Type        string
	Description string
	CreatedAt   time.Time
}

type Statement struct {
	Balance      *StatementBalance
	Transactions []*StatementTransaction
}

func (s Statement) ToJSON() []byte {
	j, _ := json.Marshal(s)
	return j
}

type TransactionService interface {
	Create(ctx *fasthttp.RequestCtx)
	GetStatement(ctx *fasthttp.RequestCtx)
}

type TransactionUseCases interface {
	Create(req *dto.CreateTransactionRequest) (*Transaction, error)
	GetStatement(accountId uint) (*Statement, error)
}

type TransactionRepository interface {
	Create(req *dto.CreateTransactionRequest) (*Transaction, error)
	GetStatement(accountId uint) (*Statement, error)
}
