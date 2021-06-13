/*
using buffered channels, it is possible to send multiple messages into a channel without being blocked
*/
package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
func main() {

	// creates capacity of 2
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(5 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		// this will take only 2 values at time, once channel flush those 2 values, it can take next 2 values
	}
}

// channel can transfer only limited data of it's capacity
