package main

import (
	"fmt"
	"sync"
)

/*
There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

The host allows no more than 2 philosophers to eat concurrently.

Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

//ChopS represents chopstick
type ChopS struct {
	sync.Mutex
}

//Philo represents philosopher
type Philo struct {
	leftCs, rightCs *ChopS
	number          int
	eatCount        int
	host            chan bool
	finished        *sync.WaitGroup
}

func (p Philo) eat() {
	for p.eatCount < 3 {
		p.host <- true //asking for permission to eat
		select {
		case permissionGranted := <-p.host:
			if permissionGranted {
				p.leftCs.Lock()
				p.rightCs.Lock()

				fmt.Println("starting to eat", p.number)
				p.eatCount++
				fmt.Println("finishing eating", p.number)

				p.rightCs.Unlock()
				p.leftCs.Unlock()
				p.host <- false //signaling to host that finished eating
			}
		}
	}
	p.finished.Done()
}

//host tracks how many concurrent philosopers to allow
func host(channels []chan bool) {
	eatersCount := 0
	for {
		select {
		case x := <-channels[0]:
			eatersCount = handleRequest(channels[0], eatersCount, x)
		case x := <-channels[1]:
			eatersCount = handleRequest(channels[1], eatersCount, x)
		case x := <-channels[2]:
			eatersCount = handleRequest(channels[2], eatersCount, x)
		case x := <-channels[3]:
			eatersCount = handleRequest(channels[3], eatersCount, x)
		case x := <-channels[4]:
			eatersCount = handleRequest(channels[4], eatersCount, x)
		}
	}
}

func handleRequest(channel chan bool, eatersCount int, x bool) int {
	if x == true && eatersCount < 2 {
		eatersCount++
		channel <- true //allow to eat
	} else if x == false {
		eatersCount-- //finished eating
	} else {
		channel <- false //signal no permission to eat, should wait more
	}
	return eatersCount
}

func main() {
	wg := new(sync.WaitGroup)
	hostChnls := make([]chan bool, 5)
	CSTicks := make([]*ChopS, 5)
	Philos := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		CSTicks[i] = new(ChopS)
	}

	go host(hostChnls)

	for i := 0; i < 5; i++ {
		hostChnls[i] = make(chan bool)
		Philos[i] = &Philo{CSTicks[i], CSTicks[(i+1)%5], i + 1, 0, hostChnls[i], wg}
		wg.Add(1)
		go Philos[i].eat()
	}
	wg.Wait()
}
