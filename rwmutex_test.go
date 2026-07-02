package belajar_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
 * RWMutex (read write mutex)
 * digunakan untuk melakukan proses locking pada saat mengubah dan membaca data yang diakses oleh goroutine
 * jadi ada dua lock, lock untuk read & write
 */

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()

	return balance
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("total balance: ", account.GetBalance())
}
