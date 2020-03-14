package main

import (
	"fmt"
	"time"
)

var start = time.Now()

func main() {
	// unbuffered()
	
	// buffered()
	
	// defaultCase()
	
	// defaultCaseWaiting()
	
	avoidDeadlock()
}

func avoidDeadlock() {
	fmt.Println("main() started", time.Since(start))
	
	chan1 := make(chan string)
	chan2 := make(chan string)
	
	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from chan2", res, time.Since(start))
	default:
		fmt.Println("No goroutines available to send data", time.Since(start))
	}
	
	fmt.Println("main() stopped", time.Since(start))
}

func defaultCaseWaiting() {
	fmt.Println("main() started", time.Since(start))
	
	chan1 := make(chan string)
	chan2 := make(chan string)
	
	go service1(chan1)
	go service2(chan2)
	
	time.Sleep(3 * time.Second)
	
	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	default: // makes select non-blocking
		fmt.Println("No response received", time.Since(start))
	}
	
	fmt.Println("main() stopped", time.Since(start))
}

func defaultCase() {
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
	default: // makes select non-blocking
		fmt.Println("No response received", time.Since(start))
	}
	
	fmt.Println("main() stopped", time.Since(start))
}

// non-blocking
func buffered() {
	fmt.Println("main() started", time.Since(start))
	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)
	
	chan1 <- "Value 1"
	chan1 <- "Value 2"
	chan2 <- "Value 1"
	chan2 <- "Value 2"
	
	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from chan2", res, time.Since(start))
	}
	
	fmt.Println("main() stopped", time.Since(start))
}

// blocking
func unbuffered() {
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
	}
	
	fmt.Println("main() stopped", time.Since(start))
}

func service1(c chan string) {
	fmt.Println("service1() started", time.Since(start))
	// time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	fmt.Println("service2() started", time.Since(start))
	// time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}
