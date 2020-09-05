/*
To wait for multiple goroutines to finish, we can use a wait group.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// our worker/go routine
func sayHii(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started running \n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d stopped running \n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go sayHii(i, &wg)
	}
	wg.Wait()
}
