package transaction_worker

import (
	"log"
	"sync"
	"time"

	"github.com/reonardoleis/banky/internal/core/dto"
)

var (
	buff     = make([]*dto.CreateTransactionRequest, 0)
	buffLock = new(sync.Mutex)
)

func (w worker) runInsert() {
	for {
		buffLock.Lock()

		if len(buff) == 0 {
			buffLock.Unlock()
			time.Sleep(1 * time.Second)
			continue
		}

		limit := 1000
		var subset []*dto.CreateTransactionRequest
		if len(buff) < limit {
			subset = buff
		} else {
			subset = buff[:limit]
		}

		err := w.usecase.Create(subset)
		if err != nil {
			log.Println("transaction worker error while creating transaction on run insert", err)
		}

		remove := len(subset)

		if remove == len(buff) {
			buff = make([]*dto.CreateTransactionRequest, 0)
		} else {
			buff = buff[remove:]
		}

		buffLock.Unlock()

		time.Sleep(1 * time.Second)
	}
}

func (w worker) Run() error {
	go w.runInsert()

	for v := range w.ch {
		buffLock.Lock()
		buff = append(buff, v)
		buffLock.Unlock()
	}

	return nil
}
