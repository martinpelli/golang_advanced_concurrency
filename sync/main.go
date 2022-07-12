package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, waitGroup *sync.WaitGroup, mutex *sync.RWMutex) {
	defer waitGroup.Done()
	mutex.RLock()
	balance += amount
	mutex.RUnlock()
}

func Balance(mutex *sync.RWMutex) int {
	mutex.RLock()
	b := balance
	mutex.RUnlock()
	return b
}

func main() {
	var waitGroup sync.WaitGroup
	//var mutex sync.Mutex
	var mutex sync.RWMutex

	for i := 1; i <= 5; i++ {
		waitGroup.Add(1)
		go Deposit(i*100, &waitGroup, &mutex)
	}

	waitGroup.Wait()
	fmt.Println(Balance(&mutex))

}
