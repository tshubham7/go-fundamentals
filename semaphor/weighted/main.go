/*
in the previous example, you can allow your go routine to run as soon as you recieve the free space
suppose you want to execute in batches,
for eg. you and your 5 friends wants to play playstation together
what you want to wait till to you get 5 free spaces so that you can play
does not matter they have 2 or 3 free space available, you will wait until you get 5.
let's see
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// play with changing the size of the array
const arraySize = 100

var numbers = [arraySize]int{}

// sem is a channel that will allow up to 10 concurrent operations.
var sem = semaphore.NewWeighted(int64(10))

func main() {
	var wg sync.WaitGroup
	for i := 0; i < arraySize; i++ {
		wg.Add(1)
		sem.Acquire(context.Background(), 1)
		go feedValue(i, &wg)
		if sem.TryAcquire(5) { //trying to occupy 5 rooms
			wg.Add(1)
			i += 5
			go feedFiveValue(i, &wg)
		}
	}
	wg.Wait()
	// fmt.Println(numbers)
}

func feedFiveValue(index int, wg *sync.WaitGroup) {
	fmt.Printf("runing process from %d to %d together\n", index, index+5)
	time.Sleep(time.Second * 2)
	for i := index; i < index+5; i++ {
		numbers[index] = (index + 1) * 2
	}

	fmt.Printf("finished process from %d to %d together\n", index, index+5)
	defer wg.Done()
	sem.Release(5)
}

// multiplying the index value by 2
func feedValue(index int, wg *sync.WaitGroup) {
	fmt.Printf("running process %d\n", index)
	time.Sleep(time.Second)

	numbers[index] = (index + 1) * 2
	defer wg.Done()

	sem.Release(1) // equivalent to <- sem (using channel approach)
}
