package main

import (
	"fmt"
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	for {
		select {
		case _, ok := <-snapshotTicker:
			if !ok {
				return
			}
			takeSnapshot()
		case _, ok := <-saveAfter:
			if !ok {
				return
			}
			saveSnapshot()
			return
		default:
			waitForData()
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// TEST SUITE - Don't touch below this line

func takeSnapshot() {
	fmt.Println("Taking a backup snapshot...")
}

func saveSnapshot() {
	fmt.Println("All backups saved!")
}

func waitForData() {
	fmt.Println("Nothing to do, waiting...")
}

func test() {
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("===========================")
}

func main() {
	test()
}
