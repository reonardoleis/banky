package domain

import (
	"time"

	"github.com/valyala/fasthttp"
)

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

type StatementService interface {
	Fetch(ctx fasthttp.RequestCtx)
}

type StatementUseCases interface {
	Fetch(userId uint) (*Statement, error)
}

type StatementRepository interface {
	Fetch(userId uint) (*Statement, error)
}
