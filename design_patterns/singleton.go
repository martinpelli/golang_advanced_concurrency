package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct {
}

func (Database) CreateSingleConnection() {
	fmt.Println("Connecting to Database")
	time.Sleep(2 * time.Second)
	fmt.Println("Database connected")
}

var db *Database
var mutex sync.Mutex

func getDatabaseIstance() *Database {
	mutex.Lock()
	defer mutex.Unlock()
	if db == nil {
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("Database already connected")
	}
	return db
}

func main() {
	var waitGroup sync.WaitGroup
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			getDatabaseIstance()
		}()
	}
	waitGroup.Wait()
}
