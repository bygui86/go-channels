package main

import (
	"fmt"
	"time"
)

func main() {
	normal(false)
	fmt.Println()
	normal(true)
	fmt.Println()
	overflowing()
}

func normal(waitAmoment bool) {
	fmt.Println("normal() started")
	c := make(chan int, 3)
	
	go squares(c)
	
	c <- 1
	c <- 2
	c <- 3
	
	if waitAmoment {
		fmt.Println("normal() waiting some seconds to give the time to other go-routines to drain out values in the channel")
		time.Sleep(5*time.Second)
	}
	
	fmt.Println("normal() stopped")
}

func overflowing() {
	fmt.Println("overflowing() started")
	c := make(chan int, 3)
	
	go squares(c)
	
	c <- 1
	c <- 2
	c <- 3
	c <- 4 // blocks here
	
	fmt.Println("overflowing() stopped")
}

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}
