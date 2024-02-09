package dto

import (
	"fmt"
	"strconv"
	"time"

	"github.com/goccy/go-json"
)

type CreateTransactionRequest struct {
	AccountId   uint      `json:"account_id"`
	Type        string    `json:"tipo"`
	Amount      int64     `json:"valor"`
	Description string    `json:"descricao"`
	Status      uint8     `json:"-"`
	Timestamp   time.Time `json:"-"`
}

func JsonToCreateTransactionRequest(body []byte) (*CreateTransactionRequest, error) {
	dto := new(CreateTransactionRequest)
	err := json.Unmarshal(body, dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (c CreateTransactionRequest) ToJSON() []byte {
	return []byte(fmt.Sprintf(
		`{"account_id":%d,"type":"%s","amount":%d,"description":"%s"}`,
		c.AccountId, c.Type, c.Amount, c.Description,
	))
}

type CreateTransactionResponse struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}

func (c CreateTransactionResponse) ToJSON() []byte {
	return []byte(fmt.Sprintf(
		`{"limite":%d,"saldo":%d}`,
		c.Limit, c.Balance,
	))
}

type WorkerRequest struct {
	Type uint8
	Data string
}

func JsonToWorkerRequest(body []byte) (*WorkerRequest, error) {
	dto := new(WorkerRequest)
	err := json.Unmarshal(body, dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (w WorkerRequest) ToCreateTransactionRequest() (*CreateTransactionRequest, error) {
	createTransactioRequest := new(CreateTransactionRequest)
	err := json.Unmarshal([]byte(w.Data), createTransactioRequest)
	return createTransactioRequest, err
}

func (w WorkerRequest) ToGetStatementRequest() (uint, error) {
	v := string(w.Data)
	id, err := strconv.Atoi(v)
	return uint(id), err
}
