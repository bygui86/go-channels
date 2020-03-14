package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("[main] started")
	
	tasks := make(chan int, 10)
	results := make(chan int, 10)
	
	fmt.Println("[main] launch worker goroutines")
	for i := 0; i < 3; i++ {
		go sqrWorker(tasks, results, i)
	}
	
	fmt.Println("[main] pass tasks to tasks-channel")
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}
	
	fmt.Printf("[main] %d running goroutines\n", runtime.NumGoroutine())
	
	fmt.Println("[main] close tasks-channel")
	close(tasks)
	
	fmt.Println("[main] receive results on result-channel from all workers")
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}
	
	fmt.Println("[main] stopped")
}

// worker than make squares
func sqrWorker(tasks <-chan int, results chan<- int, instance int) {
	fmt.Printf("[worker-%d] started\n", instance)
	
	for num := range tasks {
		time.Sleep(time.Millisecond) // simulating blocking task
		fmt.Printf("[worker %d] send result back\n", instance)
		results <- num * num
	}
	
	// once tasks-channel will be closed
	fmt.Printf("[worker-%d] stopped\n", instance)
}
