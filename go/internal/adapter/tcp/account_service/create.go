package account_service

import (
	"log"
	"net"

	"github.com/reonardoleis/banky/internal/core/dto"
)

func (s service) CreateTransaction(conn net.UDPConn, req *dto.CreateTransactionRequest) {
	limit, balance, ok, exists := s.worker.Enqueue(req)
	if !exists {
		_, err := conn.Write([]byte(ErrAccountNotFound.Error() + "\n"))
		if err != nil {
			log.Println("error while writing account not found response to conn", err)
		}

		return
	}

	response := new(dto.CreateTransactionResponse)
	response.Balance = balance
	response.Limit = limit

	responseJson := string(response.ToJSON()) + "\n"

	if !ok {
		responseJson = "." + responseJson
	}

	_, err := conn.Write([]byte(responseJson))
	if err != nil {
		log.Println("error while writing create response to conn", err)
	}
}
