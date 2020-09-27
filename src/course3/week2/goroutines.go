package main

import (
	"fmt"
	"time"
)

/**
* The goal of this activity is to explore race conditions by creating and running two simultaneous goroutines.
*
* Normally if it would be non-concurent program it would print sequentially incremented value on each iteration, but
* now results are non deterministic.
*
* It uses two go routines. Execution order of those two used go routines is non deterministic and they both depend on the same shared variable
* which is being read and altered and printed by those both routines at different points in time.
* As those two go routines execute completely independently from each other and
* program result depends on the execution order which is now non deterministic - this creates race condition or in other words non-deterministic result of the program itself.
 */

//Shared is a shared structure between two routines
type Shared struct {
	value int
}

func incrementAndPrint(counter *Shared, name string) {
	for i := 0; i < 5; i++ {
		counter.value = counter.value + 1 //this value can be altered by other routine at any time
		time.Sleep(1 * time.Millisecond)  //this make sure race-condition occurs, as routine 1 and routine 2 reads the same value before altering
		fmt.Println("Routine:", name, "value:", counter.value)
	}
}

func main() {
	for {
		fmt.Println()
		fmt.Println("I should print values from 1 to 10, but I do not!")
		fmt.Println("Press ENTER to run with race condition, repeat to observe different results each time>")
		fmt.Scanln()

		runConcurently()
	}
}

func runConcurently() {
	sharedCounter := Shared{0}

	go incrementAndPrint(&sharedCounter, "1")
	go incrementAndPrint(&sharedCounter, "2")
	time.Sleep(500 * time.Millisecond)

}
