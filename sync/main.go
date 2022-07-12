package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, waitGroup *sync.WaitGroup, lock *sync.Mutex) {
	defer waitGroup.Done()
	lock.Lock()
	balance += amount
	lock.Unlock()
}

func Balance() int {
	return balance
}

func main() {
	var waitGroup sync.WaitGroup
	var lock sync.Mutex

	for i := 1; i <= 5; i++ {
		waitGroup.Add(1)
		go Deposit(i*100, &waitGroup, &lock)
	}

	waitGroup.Wait()
	fmt.Println(Balance())

}
