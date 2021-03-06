package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main() started")
	c := make(chan string)
	
	go greet(c)
	
	c <- "John" // blocking operation as channel is unbuffered
	fmt.Println("main() stopped")
}

func greet(c chan string) {
	fmt.Println("greet() started")
	time.Sleep(3 * time.Second)
	fmt.Println("greet() Hello " + <-c + "!")
	fmt.Println("greet() stopped")
}
