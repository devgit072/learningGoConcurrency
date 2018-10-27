package main

import (
	"fmt"
	"time"
)



// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func main1() {
	go goroutine1()
	// Hello world will not be printed because main will finishes before that.
	// U can put time.Sleep() , then main will take longer time to finish and then hello world will be printed.
}

func goroutine1() {
	fmt.Println("Hello world!!")
}
// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func main2() {
	go func(message string) {
		fmt.Println(message)
	}("hello world")
	time.Sleep(1 * time.Second)
}
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func main() {
	msgPrinter := func(mesg string) {
		fmt.Println("Hello world....")
	}

	go msgPrinter("Hello world!!")
	go msgPrinter("Hello westoros")
	time.Sleep(time.Second)
}
