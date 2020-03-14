package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	
	// iteration terminates after receiving 3 values
	for elem := range c {
		fmt.Println(elem)
	}
}
