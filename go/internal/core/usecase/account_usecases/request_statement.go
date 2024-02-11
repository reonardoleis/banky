package account_usecases

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/goccy/go-json"

	"github.com/reonardoleis/banky/internal/core/domain"
	"github.com/reonardoleis/banky/internal/core/dto"
)

func (u usecase) RequestStatement(accountId uint) (*domain.Statement, error) {
	conn, err := net.Dial("tcp", os.Getenv("MANAGER_ADDR"))
	if err != nil {
		log.Println("transaction usecases error requesting statement", err)
		return nil, err
	}

	defer conn.Close()

	workerReq := &dto.WorkerRequest{
		Type: 1,
		Data: fmt.Sprintf("%d", accountId),
	}

	message, err := json.Marshal(workerReq)
	if err != nil {
		log.Println("transaction usecases error marshalling worker req before requesting statement", err)
		return nil, err
	}

	fmt.Fprintf(conn, string(message)+"\n")

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("transaction usecases error reading from conn after requesting statement", err)
		return nil, err
	}

	if response[0] == 'a' {
		return nil, nil
	}

	statement := new(domain.Statement)
	err = json.Unmarshal([]byte(response), statement)
	if err != nil {
		log.Println("transaction usecases error unmarshalling response after requesting statement", err)
		return nil, err
	}

	return statement, nil
}
