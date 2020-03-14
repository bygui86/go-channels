package main

import (
	"fmt"
	"time"
)

func main() {
	// read 10 fibonacci numbers from channel returned by `fib` function
	for fn := range startFibonacci(10) {
		fmt.Println("Current fibonacci number is", fn)
		time.Sleep(time.Second)
	}
}

// fib returns a channel which transports fibonacci numbers
func startFibonacci(length int) <-chan int {
	// make buffered channel
	c := make(chan int, length)
	
	// run generation concurrently
	go generateFibonacciSequence(c, length)
	
	// return channel
	return c
}

func generateFibonacciSequence(c chan int, length int) {
	for i, j := 0, 1; i < length; i, j = i+j, i {
		c <- i
	}
	close(c)
}
