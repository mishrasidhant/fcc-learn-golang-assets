package main

import (
	"fmt"
	"time"
)

func concurrrentFib(n int) {
	/*
		1. Create a channel
		2. Call fibonacchi in a go routine (concurrently)
		3. Run a for loop n times that checks the channel and prints
		return
	*/
	/*
		ch := make(chan int, n)
		go fibonacci(n, ch)
		for i := 0; i < n; i++ {
			fib, ok := <-ch
			if !ok {
				return
			}
			fmt.Println(fib)
		}
	*/
	ch := make(chan int, n)
	go fibonacci(n, ch)
	for fibNum := range ch {
		fmt.Println(fibNum)
	}
}

// TEST SUITE - Don't touch below this line

func test(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrrentFib(n)
	fmt.Println("==============================")
}

func main() {
	test(10)
	test(5)
	test(20)
	test(13)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
