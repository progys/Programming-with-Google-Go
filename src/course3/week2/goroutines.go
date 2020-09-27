package main

import (
	"fmt"
	"time"
)

/**
* The goal of this activity is to explore race conditions by creating and running two simultaneous goroutines.
 */

//Shared is a shard structure between two routines
type Shared struct {
	value int
}

func incrementAndPrint(counter *Shared, name string) {
	counter.value = counter.value + 1 //this value can be altered by other routine at any time
	time.Sleep(1 * time.Second)       //this guarantees that race condition occurs
	fmt.Println("Routine:", name, "value: ", counter.value)
}

/*
* Normally if it would be non-concurent program it would print icremented value on each iteration,
* now it prints the same value in each concurent iteration.
*
* It uses two go routines. Execution order of those two used go routines is non deterministic and they both depend on the same shared variable
* which is being read and altered and printed by those both routines at different points in time.
* As those two go routines execute completely independently from each other and
* program result depends on the execution order which is non deterministic - this creates race condition or in other words non-deterministic result of the program itself.
 */
func main() {
	sharedCounter := Shared{0}

	fmt.Println()
	fmt.Println("I should print values from 1 to 20. But I do not!")
	fmt.Println()
	for i := 0; i < 10; {
		go incrementAndPrint(&sharedCounter, "r1")
		go incrementAndPrint(&sharedCounter, "r2")
		i++
		time.Sleep(500 * time.Millisecond)
	}
}
