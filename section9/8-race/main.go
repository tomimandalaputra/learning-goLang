package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (b *BankAccount) Deposit(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.balance += amount
	fmt.Println("Deposit", amount)
}

func (b *BankAccount) Withdraw(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.balance < amount {
		fmt.Println("Cannot withdraw that amount:", amount)
		return
	}

	b.balance -= amount
	fmt.Println("Withdraw", amount)
}

func (b *BankAccount) Balance() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.balance
}

func ConcurrencyDeposit(amount int, ba *BankAccount, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(amount) * time.Millisecond)
	ba.Deposit(amount)
}

func ConcurrencyWithdraw(amount int, ba *BankAccount, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(amount) * time.Millisecond)
	ba.Withdraw(amount * 10)
}

func main() {
	// counter := 0 // critical section
	// var wg sync.WaitGroup
	// var mutex sync.Mutex

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)

	// 	go func() {
	// 		defer wg.Done()
	// 		mutex.Lock()
	// 		counter++
	// 		mutex.Unlock()
	// 	}()
	// }

	// wg.Wait()
	// fmt.Println("Counter", counter)

	var wg sync.WaitGroup
	var account = &BankAccount{balance: 100}

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go ConcurrencyDeposit(i+1, account, &wg)
	// }

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go ConcurrencyWithdraw(i+1, account, &wg)
	}

	wg.Wait()
	fmt.Println("Balance: ", account.Balance())
}
