package account_usecases

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/goccy/go-json"

	"github.com/reonardoleis/banky/internal/core/dto"
)

func (u usecase) DispatchTransactionCreation(req *dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, bool, error) {
	conn, err := net.Dial("tcp", os.Getenv("MANAGER_ADDR"))
	if err != nil {
		log.Println("transaction usecases error dispatching creation", err)
		return nil, false, err
	}

	defer conn.Close()

	data, err := json.Marshal(req)
	if err != nil {
		log.Println("transaction usecases error marshalling req before dispatching creation", err)
		return nil, false, err
	}

	workerReq := &dto.WorkerRequest{
		Type: 0,
		Data: string(data),
	}

	message, err := json.Marshal(workerReq)
	if err != nil {
		log.Println("transaction usecases error marshalling worker req before dispatching creation", err)
		return nil, false, err
	}

	fmt.Fprintf(conn, string(message)+"\n")

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("transaction usecases error reading from conn after dispatching creation", err)
		return nil, false, err
	}

	if response[0] == 'a' {
		return nil, false, nil
	}

	ok := !(response[0] == '.')
	if !ok {
		response = response[1:]
	}

	transaction := new(dto.CreateTransactionResponse)
	err = json.Unmarshal([]byte(response), transaction)
	if err != nil {
		log.Println("transaction usecases error unmarshalling response after dispatching creation", err)
		return nil, false, err
	}

	return transaction, ok, nil
}
