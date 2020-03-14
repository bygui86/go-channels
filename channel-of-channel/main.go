package main

import "fmt"

func main() {
	fmt.Println("main() started")
	
	// make a channel `cc` of data type channel of string data type
	cc := make(chan chan string)
	
	go greeter(cc) // start `greeter` goroutine using `cc` channel
	
	// receive a channel `c` from `greeter` goroutine
	c := <-cc
	
	go greet(c) // start `greet` goroutine using `c` channel
	
	// send data to `c` channel
	c <- "John"
	
	fmt.Println("main() stopped")
}

// gets a channels and writes a channel to it
func greeter(cc chan chan string) {
	c := make(chan string)
	cc <- c
}

// gets a channel and prints the greeting by reading from channel
func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}
