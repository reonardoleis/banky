package dto

import "encoding/json"

type CreateTransactionRequest struct {
	AccountId   uint   `json:"account_id"`
	Type        string `json:"type"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}

func JsonToCreateTransactionRequest(body []byte) (*CreateTransactionRequest, error) {
	dto := new(CreateTransactionRequest)
	err := json.Unmarshal(body, dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}
