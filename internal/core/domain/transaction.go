package domain

import (
	"fmt"
	"net"
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
	AccountId   uint
	Status      uint8
}

func (t Transaction) ToJSON() []byte {
	return []byte(fmt.Sprintf(
		`{"valor":%d,"tipo":"%s","descricao":"%s"}`,
		t.Amount,
		t.Type,
		t.Description,
	))
}

type StatementBalance struct {
	Total int64
	Limit int64
	At    time.Time
}

func (s StatementBalance) ToJSON() []byte {
	return []byte(fmt.Sprintf(
		`{"total":%d,"limite":%d,"data_extrato":"%s"}`,
		s.Total, s.Limit, s.At.UTC().Format("2006-01-02T15:04:05.000000Z"),
	))
}

type StatementTransaction struct {
	Amount      int64
	Type        string
	Description string
	CreatedAt   time.Time
}

func (s StatementTransaction) ToJSON() []byte {
	return []byte(fmt.Sprintf(
		`{"valor":%d,"tipo":"%s","descricao":"%s","realizada_em":"%s"}`,
		s.Amount, s.Type, s.Description, s.CreatedAt.UTC().Format("2006-01-02T15:04:05.000000Z"),
	))
}

type Statement struct {
	Balance      *StatementBalance
	Transactions []*StatementTransaction
}

func (s Statement) ToJSON() []byte {
	transactions := "["
	for i, t := range s.Transactions {
		transactions += string(t.ToJSON())
		if i < len(s.Transactions)-1 {
			transactions += ","
		}
	}
	transactions += "]"

	balance := string(s.Balance.ToJSON())

	return []byte(fmt.Sprintf(
		`{"saldo":%s,"ultimas_transacoes":%s}`,
		balance, transactions,
	))
}

type TransactionApiService interface {
	DispatchCreation(ctx *fasthttp.RequestCtx)
	RequestStatement(ctx *fasthttp.RequestCtx)
}

type TransactionManagerService interface {
	Create(conn net.Conn, req *dto.CreateTransactionRequest)
	GetStatement(conn net.Conn, accountId uint)
	Worker() TransactionWorker
}

type AccountInformation struct {
	Account          *Account
	LastTransactions []*Transaction
	Balance          int64
	Limit            int64
}

func (a *AccountInformation) Push(t *Transaction) {
	if len(a.LastTransactions) == 10 {
		a.LastTransactions = a.LastTransactions[1:]
	}

	a.LastTransactions = append(a.LastTransactions, t)
}

type TransactionUseCases interface {
	Create(req []*dto.CreateTransactionRequest) error
	GetStatement(accountId uint, accounts map[uint]*AccountInformation) (*Statement, bool, error)
	DispatchCreation(req *dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, bool, error)
	RequestStatement(accountId uint) (*Statement, error)
}

type TransactionRepository interface {
	Create(req []*dto.CreateTransactionRequest) error
	GetStatement(accountId uint) (*Statement, error)
	GetLastTransactions(accountId uint) ([]*Transaction, error)
}

type TransactionWorker interface {
	Enqueue(req *dto.CreateTransactionRequest) (limit int64, balance int64, ok bool, exists bool)
	GetStatement(accountId uint) (statement *Statement, exists bool, err error)
	Run() error
}
