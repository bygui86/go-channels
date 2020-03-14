package main

import "fmt"

func main() {
	fmt.Println("main() started")
	c := make(chan int)
	
	go squares(c) // start goroutine
	
	// periodic block/unblock of main goroutine until channel closes
	for val := range c {
		fmt.Println(val)
	}
	
	fmt.Println("main() stopped")
}

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}
	
	close(c) // close channel
}
