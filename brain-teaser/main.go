package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)
	go squares(c)
	
	fmt.Println("active goroutines", runtime.NumGoroutine())
	c <- 1
	c <- 2
	c <- 3
	c <- 4 // blocks here
	
	fmt.Println("active goroutines", runtime.NumGoroutine())
	
	go squares(c)
	
	fmt.Println("active goroutines", runtime.NumGoroutine())
	
	c <- 5
	c <- 6
	c <- 7
	c <- 8 // blocks here
	
	fmt.Println("active goroutines", runtime.NumGoroutine())
	fmt.Println("main() stopped")
}

func squares(c chan int) {
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}
