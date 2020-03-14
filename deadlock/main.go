package main

import "fmt"

func main() {
	fmt.Println("main() started")
	
	c := make(chan string)
	c <- "John"
	
	fmt.Println("main() stopped")
}
