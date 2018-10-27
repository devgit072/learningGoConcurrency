package main

import (
	"fmt"
	"sync"
)

func main() {
	manyGoRoutines()
	fmt.Println("Exiting main now....")
}

func waitGroupDemo() {
	var wait sync.WaitGroup
	wait.Add(1) // wait for 1 go routine, i.e put 1 value in queue
	go func() {
		fmt.Println("Yoyo..")
		fmt.Println("Lalallala")
		wait.Done() // it says that tell waitGroup that 1 go routine finished, so waitGroup will minus 1 in queue
	}()
	wait.Wait() // it says that halt in this point until value in wait queue becomes 0.
}

func waitGroupDemo1() {
	var wait sync.WaitGroup
	wait.Add(1) // wait for 1 go routine, i.e put 1 value in queue
	go func() {
		fmt.Println("Kapi coffeee..")
		fmt.Println("Masala tea...")
		wait.Add(-1) // wait.Done() decreses one from queue, so Add(-1) also do same thing only... both are same.
	}()
	wait.Wait() // it says that halt in this point until value in wait queue becomes 0.
}


//lets create 5 parallel goroutine to print something.....
func manyGoRoutines() {
	var waitGroup sync.WaitGroup
	numberOfGoroutines := 10
	waitGroup.Add(numberOfGoroutines)

	for i:= 0 ;i< numberOfGoroutines ;i++ {
		go func(goRoutineId int) {
			fmt.Println("Hello goroutine: ", goRoutineId)
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()// wait till count in queue becomes 0.
}