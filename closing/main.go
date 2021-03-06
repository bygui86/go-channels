package main

import "fmt"

func main() {
	fmt.Println("main() started")
	
	c := make(chan string, 1)
	
	go greet(c)
	c <- "John"
	
	close(c) // closing channel
	
	c <- "Mike"
	fmt.Println("main() stopped")
}

func greet(c chan string) {
	<-c // for John
	<-c // for Mike
}
