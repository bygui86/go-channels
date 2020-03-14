package main

import (
	"fmt"
	"time"
)

var start = time.Now()

func main() {
	fmt.Println("main() started", time.Since(start))
	
	chan1 := make(chan string)
	chan2 := make(chan string)
	
	go service1(chan1)
	go service2(chan2)
	
	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	case <-time.After(2 * time.Second):
		fmt.Println("No response received", time.Since(start))
	}
	
	fmt.Println("main() stopped", time.Since(start))
}

func service1(c chan string) {
	fmt.Println("service1() started", time.Since(start))
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	fmt.Println("service2() started", time.Since(start))
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}
