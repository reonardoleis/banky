package account_worker

import (
	"log"
	"sync"

	"github.com/reonardoleis/banky/internal/core/domain"
	"github.com/reonardoleis/banky/internal/core/dto"
)

var accounts = make(map[uint]*domain.AccountInformation)

type worker struct {
	locks   map[uint]*sync.Mutex
	ch      chan *dto.CreateTransactionRequest
	usecase domain.AccountUseCases
}

func Initalize(
	accountRepository domain.AccountRepository,
) error {
	var err error
	accounts, err = accountRepository.LoadAccounts()
	if err != nil {
		log.Println("transaction worker error while initializing accounts", err)
		return err
	}

	for _, account := range accounts {
		account.LastTransactions, err = accountRepository.GetLastTransactions(account.Account.ID)
		if err != nil {
			log.Println("transaction worker error while initializing transactions", err)
			return err
		}
	}

	return nil
}

func New(usecase domain.AccountUseCases) domain.AccountWorker {
	locks := make(map[uint]*sync.Mutex)
	for k := range accounts {
		locks[k] = new(sync.Mutex)
	}

	return &worker{
		locks:   locks,
		ch:      make(chan *dto.CreateTransactionRequest),
		usecase: usecase,
	}
}
