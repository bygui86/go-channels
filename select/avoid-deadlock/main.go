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
