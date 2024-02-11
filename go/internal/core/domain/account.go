package domain

import (
	"fmt"
	"net"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/reonardoleis/banky/internal/core/dto"
)

type Account struct {
	ID             uint
	InitialBalance int64
	Limit          int64
}

type Transaction struct {
	ID          uint
	Type        string
	Amount      int64
	Description string
	CreatedAt   time.Time
	AccountId   uint
	Status      uint8
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

type AccountApiService interface {
	DispatchTransactionCreation(ctx *fasthttp.RequestCtx)
	RequestStatement(ctx *fasthttp.RequestCtx)
}

type AccountManagerService interface {
	CreateTransaction(conn net.Conn, req *dto.CreateTransactionRequest)
	GetStatement(conn net.Conn, accountId uint)
	Worker() AccountWorker
}

type AccountInformation struct {
	Account          *Account
	LastTransactions []*Transaction
	Balance          int64
	Limit            int64
}

type AccountUseCases interface {
	HandleTransaction(
		accountInformation *AccountInformation,
		req *dto.CreateTransactionRequest,
	) (limit, balance int64, ok bool)
	CreateTransaction(req []*dto.CreateTransactionRequest) error
	GetStatement(accountId uint, accounts map[uint]*AccountInformation) (*Statement, bool, error)
	DispatchTransactionCreation(
		req *dto.CreateTransactionRequest,
	) (*dto.CreateTransactionResponse, bool, error)
	RequestStatement(accountId uint) (*Statement, error)
}

type AccountRepository interface {
	CreateTransaction(req []*dto.CreateTransactionRequest) error
	GetStatement(accountId uint) (*Statement, error)
	GetLastTransactions(accountId uint) ([]*Transaction, error)
	LoadAccounts() (map[uint]*AccountInformation, error)
}

type AccountWorker interface {
	Enqueue(req *dto.CreateTransactionRequest) (limit int64, balance int64, ok bool, exists bool)
	GetStatement(accountId uint) (statement *Statement, exists bool, err error)
	Run() error
}

// JSON/parsers to entities
func (a *AccountInformation) Push(t *Transaction) {
	if len(a.LastTransactions) == 10 {
		a.LastTransactions = a.LastTransactions[1:]
	}

	a.LastTransactions = append(a.LastTransactions, t)
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

func (s StatementTransaction) ToJSON() []byte {
	return []byte(fmt.Sprintf(
		`{"valor":%d,"tipo":"%s","descricao":"%s","realizada_em":"%s"}`,
		s.Amount, s.Type, s.Description, s.CreatedAt.UTC().Format("2006-01-02T15:04:05.000000Z"),
	))
}

func (s StatementBalance) ToJSON() []byte {
	return []byte(fmt.Sprintf(
		`{"total":%d,"limite":%d,"data_extrato":"%s"}`,
		s.Total, s.Limit, s.At.UTC().Format("2006-01-02T15:04:05.000000Z"),
	))
}

func (t Transaction) ToJSON() []byte {
	return []byte(fmt.Sprintf(
		`{"valor":%d,"tipo":"%s","descricao":"%s"}`,
		t.Amount,
		t.Type,
		t.Description,
	))
}
