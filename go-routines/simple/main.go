/*A Goroutine is a function or method which executes independently
and simultaneously in connection with any other Goroutines present in your program.
you are allowed to create multiple goroutines in a single program.
You can create a goroutine simply by using go keyword as a prefixing to the function
or method call as shown in the below
source: geeksforgeeks*/
package main

import (
	"fmt"
	"time"
)

// go routine 1
func printNames() {
	arr1 := [4]string{"Pallu", "Sallu", "Ballu", "Kallu"}

	for t1 := 0; t1 <= 3; t1++ {

		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%s\n", arr1[t1])
	}
}

// goroutine 2
func sayHii() {

	for i := 0; i <= 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("hii...")
	}
}

func main() {

	fmt.Println("!...Main Go-routine Start...!")

	// calling Goroutine 1
	go printNames()

	// calling Goroutine 2
	go sayHii()

	time.Sleep(3500 * time.Millisecond)
	fmt.Println("\n!...Main Go-routine End...!")
}

/* output
!...Main Go-routine Start...!
Pallu
Sallu
Ballu
hii...
Kallu
hii...
hii...
hii...

!...Main Go-routine End...!
*/
