package handler

import (
	"bufio"
	"log"
	"net"

	"github.com/reonardoleis/banky/internal/core/domain"
	"github.com/reonardoleis/banky/internal/core/dto"
	"golang.org/x/sync/semaphore"
)

func Handle(s domain.TransactionManagerService, conn net.Conn, sem *semaphore.Weighted) {
	defer conn.Close()
	defer sem.Release(1)
	message, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Println("error while handling message", err)
		return
	}

	workerRequest, err := dto.JsonToWorkerRequest(message)
	if err != nil {
		log.Println("error while parsing worker request", err)
		_, err := conn.Write([]byte(err.Error() + "\n"))
		if err != nil {
			log.Println("error while sending error response", err)
		}

		return
	}

	if workerRequest.Type == 0 {
		req, err := workerRequest.ToCreateTransactionRequest()
		if err != nil {
			log.Println("error while parsing worker request to create transaction request", err)
			_, err := conn.Write([]byte(err.Error() + "\n"))
			if err != nil {
				log.Println("error while sending error response", err)
			}

			return
		}

		s.Create(conn, req)
		return
	} else {
		req, err := workerRequest.ToGetStatementRequest()
		if err != nil {
			log.Println("error while parsing worker request to get statement", err)
			_, err := conn.Write([]byte(err.Error() + "\n"))
			if err != nil {
				log.Println("error while sending error response", err)
			}
			return
		}
		s.GetStatement(conn, req)
	}
}
