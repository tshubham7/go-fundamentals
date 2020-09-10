/*
I am considering that you are aware of mutex, if not check out the mutex example
Mutex you can lock a resource so that only current transaction or goroutine can access it.

Semaphor is similar to mutex but you allow multiple go routines  to access your resource.
You can actually use buffered channels to limit the amount of go routines for access.
let's check out
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// play with changing the size of the array
const arraySize = 100

var numbers = [arraySize]int{}

// sem is a channel that will allow up to 10 concurrent operations.
var sem = make(chan int, 10)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < arraySize; i++ {
		wg.Add(1)
		sem <- 1
		go feedValue(i, &wg)
	}
	wg.Wait()
	fmt.Println(numbers)
}

// multiplying the index value by 2
func feedValue(index int, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	numbers[index] = (index + 1) * 2
	fmt.Printf("running for index %d\n", index)
	defer wg.Done()
	<-sem
}
