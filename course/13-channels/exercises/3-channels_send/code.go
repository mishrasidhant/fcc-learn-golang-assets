package main

import (
	"fmt"
	"time"
)

func waitForDbs(numDBs int, dbChan chan struct{}) {
	// go func() {
	for i := 0; i < numDBs; i++ {
		time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
		<-dbChan
		fmt.Printf("Waiting for db number %v", i+1)
	}
	// }()
}

// don't touch below this line

func test(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	// time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
	time.Sleep(time.Millisecond * 10000) // ensure the last print statement happens
	fmt.Println("All databases are online!")
	fmt.Println("=====================================")
}

func main() {
	test(3)
	test(4)
	test(5)
}

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			time.Sleep(time.Millisecond * 100)
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()
	return ch
}
