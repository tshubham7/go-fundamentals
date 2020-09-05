/*
when you use channels with multiple go routines
one way you can do this by using channels
*/
package main

import (
	"fmt"
	"sync"
)

func getDouble(ch chan int, value int) {
	defer wg.Done()
	ch <- value * 2
}

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(1) // telling that we are adding 1 go routine
		go getDouble(ch, i)
	}
	wg.Wait()
	close(ch)

	for value := range ch {
		fmt.Println(value)
	}
}
