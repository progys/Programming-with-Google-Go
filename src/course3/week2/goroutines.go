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
	doubleSum int
}

func doubleAndPrint(x *Shared, name string) {
	x.doubleSum = x.doubleSum + x.doubleSum //race condition here, because this value can be altered by other routine at any time
	fmt.Println("Routine:", name, "sum: ", x.doubleSum)
}

/*
* Normally if it would non-concurent program it would print doubling values on each iteration,
* now it sometimes prints the same value in each concurent iteration or prints the smaller value after larger one.
* Because execution order is not deterministic and this program depends on same shared value "sharedSum" which is being altered in non-deterministic way.
 */
func main() {
	sharedSum := Shared{1}

	for i := 0; i < 6; {
		go doubleAndPrint(&sharedSum, "r1")
		go doubleAndPrint(&sharedSum, "r2")
		time.Sleep(1 * time.Second)
		i++
	}
}
