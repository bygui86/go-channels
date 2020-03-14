package main

import (
	"fmt"
	"sync"
)

// fan-in & fan-out
func main() {
	// step 1: get input numbers channel
	// by calling `getGeneratorChannel` function, it runs a goroutine which sends number to returned channel
	numbersChannel := getGeneratorChannel()
	
	// step 2: `fan-out` square operations to multiple goroutines
	// this can be done by calling `getSquareChannel` function multiple times where individual function call returns a channel which sends square of numbers provided by `numbersChannel` channel
	// `getSquareChannel` function runs goroutines internally where squaring operation is ran concurrently
	sqrChannel1 := getSquareChannel(numbersChannel)
	sqrChannel2 := getSquareChannel(numbersChannel)
	
	// step 3: `fan-in` (combine) `sqrChannel1` and `sqrChannel2` output to merged channel
	// this is achieved by calling `channelMerger` function which takes multiple channels as arguments
	// and using `WaitGroup` and multiple goroutines to receive square number, we can send square numbers
	// to `merged` channel and close it
	mergeChannel := channelMerger(sqrChannel1, sqrChannel2)
	
	// step 4: let's sum all the squares from 0 to 9 which should be about `285`
	// this is done by using `for range` loop on `mergeChannel`
	sqrSum := 0
	
	// run until `mergeChannel` or merged channel closes
	// that happens in `channelMerger` function when all goroutines pushing to merged channel finishes
	// check line no. 86 and 87
	for sqr := range mergeChannel {
		sqrSum += sqr
	}
	
	// step 5: print sum when above `for loop` is done executing which is after `mergeChannel` channel closes
	fmt.Println("Sum of squares between 0-9 is", sqrSum)
}

// return channel for input numbers
func getGeneratorChannel() <-chan int {
	// make return channel
	input := make(chan int, 100)
	
	// run goroutine
	go numberGenerator(input)
	
	// return channel
	return input
}

func numberGenerator(input chan int) {
	// close channel once all numbers are sent to channel
	defer close(input)
	
	for num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		input <- num
	}
}

// returns a channel which returns square of numbers
func getSquareChannel(input <-chan int) <-chan int {
	// make return channel
	output := make(chan int, 100)
	
	// run goroutine
	go squareCalculator(input, output)
	
	// return channel
	return output
}

func squareCalculator(input <-chan int, output chan<- int) {
	// close output channel once for loop finishes
	defer close(output)
	
	// push squares until input channel closes
	for num := range input {
		output <- num * num
	}
}

// returns a merged channel of `outputsChan` channels
// this produce fan-in channel
// this a variadic function, meaning that takes a slice of channels as input
func channelMerger(outputsChan ...<-chan int) <-chan int {
	// create a WaitGroup
	var wg sync.WaitGroup
	
	// make return channel
	merged := make(chan int, 100)
	
	// increase counter to number of channels `len(outputsChan)`
	// as we will spawn number of goroutines equal to number of channels received to channelMerger
	wg.Add(len(outputsChan))
	
	// run above `output` function as goroutines, `n` number of times
	// where n is equal to number of channels received as argument the function
	// here we are using `for range` loop on `outputsChan` hence no need to manually tell `n`
	for _, optChan := range outputsChan {
		go channelBridger(optChan, merged, &wg)
	}
	
	// run goroutine to close merged channel once done
	go closeMergeChannel(&wg, merged)
	
	// return channel
	return merged
}

// channelBridger accepts an input channel (which sends square numbers), an output channel (to push numbers to)
// and a WaitGroup to let the caller know when process is completed
func channelBridger(input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	// once channel (square numbers sender) closes,
	// call `Done` on `WaitGroup` to decrement counter
	defer wg.Done()
	
	// run until channel (square numbers sender) closes
	for value := range input {
		output <- value
	}
}

func closeMergeChannel(wg *sync.WaitGroup, merged chan int) {
	// wait until WaitGroup finishes
	wg.Wait()
	close(merged)
}
