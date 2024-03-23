package singleton

import (
	"fmt"
	"sync"
)

type DB struct {
}

var db *DB

var lock = &sync.Mutex{}

func GetInstance() *DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			fmt.Println("Creating DB instance now")
			db = &DB{}
			return db
		} else {
			fmt.Println("DB instance already created inner")
			return db
		}

	} else {
		fmt.Println("DB instance already created outer")
		return db
	}
}
