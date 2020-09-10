/*
this enables you lock the resource, mutual exclusive lock
you have heard about shared lock and exclusive lock in the Database.
this is similar concept to it, you are just locking the resource until you are using it.
ref: https://godoc.org/sync#Mutex
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex
	balance int
)

/*consider an example of bank account where you account is processing through multiple
transaction at the same time
in order to avoid 'dirty read' see more -> https://www.geeksforgeeks.org/dbms-dirty-read-in-sql/
while processing a transaction, we will lock our resource until the transaction is done with it

be careful when using it, it can lead to deadlock as your resource might not unlock due to lack of
attention, make sure you always unlock the resource after you r done with it

to see the difference, try to comment the mutex.Lock() and mutex.Unlock() line in both go routines
*/
func main() {
	balance = 2000

	var wg sync.WaitGroup
	wg.Add(2) // 2 go routines

	go withdraw(500, &wg)
	go deposit(700, &wg)
	wg.Wait()

	fmt.Printf("total balance: %d\n", balance)
}

func withdraw(amount int, wg *sync.WaitGroup) {
	mutex.Lock()           // locking the resource
	tempBalance := balance // reading value

	fmt.Printf("Withdrawing %d amount from total balance: %d\n", amount, balance)
	time.Sleep(time.Second)
	balance = tempBalance - amount

	defer mutex.Unlock() // we have to unlock it anyhow
	defer wg.Done()
}

func deposit(amount int, wg *sync.WaitGroup) {
	mutex.Lock()           // locking the resource
	tempBalance := balance // reading value

	fmt.Printf("Depositing %d amount to total balance: %d\n", amount, balance)
	time.Sleep(time.Second)
	balance = tempBalance + amount

	defer mutex.Unlock() // we have to unlock it anyhow
	defer wg.Done()
}
