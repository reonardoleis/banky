package handler

import (
	"log"
	"net"

	"github.com/reonardoleis/banky/internal/core/domain"
	"github.com/reonardoleis/banky/internal/core/dto"
)

func Handle(s domain.AccountManagerService, message [1024]byte, conn *net.UDPConn, addr *net.UDPAddr) {
	workerRequest, err := dto.JsonToWorkerRequest([]byte(message[:]))
	if err != nil {
		log.Println("error while parsing worker request", err)
		_, err := conn.WriteToUDP([]byte(err.Error()+"\n"), addr)
		if err != nil {
			log.Println("error while sending error response", err)
		}

		return
	}

	if workerRequest.Type == 0 {
		req, err := workerRequest.ToCreateTransactionRequest()
		if err != nil {
			log.Println("error while parsing worker request to create transaction request", err)
			_, err := conn.WriteToUDP([]byte(err.Error()+"\n"), addr)
			if err != nil {
				log.Println("error while sending error response", err)
			}

			return
		}

		s.CreateTransaction(conn, req)
		return
	} else {
		req, err := workerRequest.ToGetStatementRequest()
		if err != nil {
			log.Println("error while parsing worker request to get statement", err)
			_, err := conn.WriteToUDP([]byte(err.Error()+"\n"), addr)
			if err != nil {
				log.Println("error while sending error response", err)
			}
			return
		}
		s.GetStatement(conn, req)
	}
}
