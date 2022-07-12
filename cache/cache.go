package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFibonacci(n int) int {
	fmt.Printf("Calculate Expensive Fibonacci for %d\n", n)
	time.Sleep(5 * time.Second)
	return n
}

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Mutex      sync.RWMutex
}

func (service *Service) Work(job int) {
	service.Mutex.RLocker()
	exists := service.InProgress[job]
	if exists {
		service.Mutex.RUnlock()
		response := make(chan int)
		defer close(response)

		service.Mutex.Lock()
		service.IsPending[job] = append(service.IsPending[job], response)
		service.Mutex.Unlock()
		fmt.Printf("Waiting for response job: %d\n", job)
		resp := <-response
		fmt.Printf("Response Done, received %d\n", resp)
		return
	}
	service.Mutex.RUnlock()

	service.Mutex.Lock()
	service.InProgress[job] = true
	service.Mutex.Unlock()

	fmt.Printf("Calculating Fibonacci for %d\n", job)
	result := ExpensiveFibonacci(job)

	service.Mutex.RLock()
	pendingWorkers, exists := service.IsPending[job]
	service.Mutex.RUnlock()

	if exists {
		for _, penpendingWorker := range pendingWorkers {
			penpendingWorker <- result
		}
		fmt.Printf("Result sent - all pending workers ready job:%d\n", job)
	}

	service.Mutex.Lock()
	service.InProgress[job] = false
	service.IsPending[job] = make([]chan int, 0)
	service.Mutex.Unlock()
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func main() {
	service := NewService()
	jobs := []int{3, 4, 5, 5, 4, 8, 8, 8}
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(jobs))
	for _, n := range jobs {
		go func(job int) {
			defer waitGroup.Done()
			service.Work(job)
		}(n)
	}
	waitGroup.Wait()
}
