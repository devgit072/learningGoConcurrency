package main

import (
	"fmt"
	"sync"
	"time"
)

// Mutext is nothing but who give permission to several goroutines to access common variable.
type Counter struct {
	sync.Mutex// Note than no mut sync.Mutex , only type needed.
    counter int
}

// You can run the command: go run -race mutexExample.go
// Above will throw exception if there is potential race issue.
// What is race condition BTW?
// When several threads involved in common variable and one of those are involved in write.
// When all threads are involved in read only then there is no race condition involved.

func main() {
	counter := Counter{}

	for i:=0 ;i< 10 ;i++ {
		go func(val int) {
			fmt.Println("Here is value:", val)
			// Since all goroutines are accessing same value, there must be some locking mechanisms.
			counter.Lock() // It says that mutex has locked this variable counter. Other goroutine will wait now.
			counter.counter++ // This is to made sure that only one goroutine is incrementine the count at one time.
			defer counter.Unlock() // Unlock the common variable.
		}(i)
	}
	time.Sleep(time.Second)
	//since several goroutines will be using it cocorrently, so it is
	// needed to lock this variable while reading.
	counter.Lock()
	defer counter.Unlock()
	fmt.Println("Counter value should be 10, what is value then: ", counter.counter)
}

// go run -race mutexExample.go
// will throw race issue for this main program.
func main1() {
	counter := Counter{}

	for i:=0 ;i< 10 ;i++ {
		go func(val int) {
			fmt.Println("Here is value:", val)
			counter.counter++
		}(i)
	}
	time.Sleep(time.Second)
	counter.Lock()
	defer counter.Unlock()
	fmt.Println("Counter value should be 10, what is value then: ", counter.counter)
}
