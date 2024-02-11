package account_service

import (
	"log"
	"net"

	"github.com/goccy/go-json"
)

func (s service) GetStatement(conn net.Conn, accountId uint) {
	statement, exists, err := s.worker.GetStatement(accountId)
	if err != nil {
		_, err = conn.Write([]byte(err.Error() + "\n"))
		if err != nil {
			log.Println("error while sending get statement error response", err)
		}
		return
	}

	if !exists {
		_, err = conn.Write([]byte(ErrAccountNotFound.Error() + "\n"))
		if err != nil {
			log.Println("error while sending get statement account not found response", ErrAccountNotFound)
		}

		return
	}

	statementJson, err := json.Marshal(statement)
	if err != nil {
		_, err = conn.Write([]byte(err.Error() + "\n"))
		if err != nil {
			log.Println("error while sending get statement error response", err)
		}
		return
	}

	responseJson := string(statementJson) + "\n"

	_, err = conn.Write([]byte(responseJson))
	if err != nil {
		log.Println("error while sending get statement response", err)
	}
}
