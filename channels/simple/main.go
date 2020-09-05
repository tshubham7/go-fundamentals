/*channels are way to interact with your go routine
you can use it to stop your go routine or can recieve values from it
*/
package main

import "fmt"

func getDoubleValue(ch chan int, v int) {
	ch <- v * 2 // send opertor that actually send values to the channel
}

func main() {
	val := make(chan int)
	go getDoubleValue(val, 5)
	result := <-val // receiving value from channel
	fmt.Println(result)
}
