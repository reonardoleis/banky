package tcp

import (
	"log"
	"net"

	"github.com/reonardoleis/banky/internal/adapter/postgres"
	"github.com/reonardoleis/banky/internal/adapter/postgres/account_repository"
	account_worker "github.com/reonardoleis/banky/internal/adapter/tcp/account_worker"
	"github.com/reonardoleis/banky/internal/adapter/tcp/handler"
	"github.com/reonardoleis/banky/internal/di"
)

func Run() error {
	err := account_worker.Initalize(account_repository.New(postgres.DB()))
	if err != nil {
		log.Fatalln("error while initializing transaction worker", err)
	}

	service := di.AccountManager(postgres.DB())

	go service.Worker().Run()

	addr, err := net.ResolveUDPAddr("udp", "65000")
	if err != nil {
		log.Fatalln("error while resolving udp address", err)
	}

	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln("error while listening on port 65000", err)
	}

	for {
		var buf [1024]byte
		_, addr, err := listener.ReadFromUDP(buf[0:])
		if err != nil {
			log.Println("error while reading from udp", err)
			continue
		}

		go handler.Handle(service, buf, listener, addr)
	}
}
