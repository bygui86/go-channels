package main

import "fmt"

func main() {
	fmt.Println("[main] started")
	
	squareChan := make(chan int)
	cubeChan := make(chan int)
	
	go square(squareChan)
	go cube(cubeChan)
	
	testNum := 3
	fmt.Println("[main] sent testNum to squareChan")
	
	squareChan <- testNum
	
	fmt.Println("[main] resuming")
	fmt.Println("[main] sent testNum to cubeChan")
	
	cubeChan <- testNum
	
	fmt.Println("[main] resuming")
	fmt.Println("[main] reading from channels")
	
	squareVal, cubeVal := <-squareChan, <-cubeChan
	sum := squareVal + cubeVal
	
	fmt.Println("[main] sum of square and cube of", testNum, "is", sum)
	fmt.Println("[main] stopped")
}

func square(c chan int) {
	fmt.Println("[square] started")
	num := <-c
	fmt.Println("[square] resuming")
	c <- num * num
	fmt.Println("[square] stopped")
}

func cube(c chan int) {
	fmt.Println("[cube] started")
	num := <-c
	fmt.Println("[cube] resuming")
	c <- num * num * num
	fmt.Println("[cube] stopped")
}
