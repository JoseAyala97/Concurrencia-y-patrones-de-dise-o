package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

// Metodo del struct Database
func (Database) CreateSingleConnection() {
	fmt.Println("Creating Singleton for Database")
	time.Sleep(5 * time.Second)
	fmt.Println("Creation Done")
}

var db *Database
var lock sync.Mutex

func getDatabaseIntance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating DB Connection")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB already created")
	}
	return db
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			// defer Se ejecuta al final de toda la ejecución
			defer wg.Done()
			getDatabaseIntance()
		}()
	}
	wg.Wait()
}
