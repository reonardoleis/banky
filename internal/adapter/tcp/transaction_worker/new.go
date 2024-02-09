package transaction_worker

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
	usecase domain.TransactionUseCases
}

func Initalize(
	accountRepository domain.AccountRepository,
	transactionRepository domain.TransactionRepository,
) error {
	var err error
	accounts, err = accountRepository.LoadAccounts()
	if err != nil {
		log.Println("transaction worker error while initializing accounts", err)
		return err
	}

	for _, account := range accounts {
		account.LastTransactions, err = transactionRepository.GetLastTransactions(account.Account.ID)
		if err != nil {
			log.Println("transaction worker error while initializing transactions", err)
			return err
		}
	}

	return nil
}

func New(usecase domain.TransactionUseCases) domain.TransactionWorker {
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
