package tcp

import (
	"context"
	"log"
	"net"

	"github.com/reonardoleis/banky/internal/adapter/postgres"
	"github.com/reonardoleis/banky/internal/adapter/postgres/account_repository"
	account_worker "github.com/reonardoleis/banky/internal/adapter/tcp/account_worker"
	"github.com/reonardoleis/banky/internal/adapter/tcp/handler"
	"github.com/reonardoleis/banky/internal/di"
	"golang.org/x/sync/semaphore"
)

func Run() error {
	err := account_worker.Initalize(account_repository.New(postgres.DB()))
	if err != nil {
		log.Fatalln("error while initializing transaction worker", err)
	}

	service := di.AccountManager(postgres.DB())

	go service.Worker().Run()

	listener, err := net.Listen("tcp", ":65000")
	if err != nil {
		log.Fatalln("error while listening on port 65000", err)
	}

	sem := semaphore.NewWeighted(50_000)
	for {
		sem.Acquire(context.Background(), 1)
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error while accepting connection", err)
			continue
		}

		go handler.Handle(service, conn, sem)
	}
}
