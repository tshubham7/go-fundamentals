/*channels are way to interact with your go routine
you can use it to stop your go routine or can recieve values from it
*/
package main

import (
	"fmt"
	"time"
)

func getDoubleValue(ch chan int, v int) {
	ch <- v * 2 // send opertor that actually send values to the channel
}

func main() {
	do1()
}

func do() {
	val := make(chan int)
	go getDoubleValue(val, 5)
	result := <-val // receiving value from channel
	fmt.Println(result)
}

// using seleect
func do1() {
	msg1 := make(chan string)
	msg2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second / 2)
			msg1 <- "faster one"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			msg2 <- "slower one"
		}
	}()

	// 	// using this approach you won't be able to print msgs in correct order
	// 	// because msg2 is blocking msg1
	// 	for {
	// 		fmt.Println(<-msg1)
	// 		fmt.Println(<-msg2)
	// 	}

	// this won't block
	for {
		select {
		case msg := <-msg1:
			fmt.Println(msg)
		case msg := <-msg2:
			fmt.Println(msg)
		}
	}

}
