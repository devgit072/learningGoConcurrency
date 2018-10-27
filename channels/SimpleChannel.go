package main

import "fmt"

// Remember channel get and take is blocking call.
// val <- channel is blockin and it will not move forward until <- returns value.
//Similarly channel <- val is also blocking

func main1() {
	//ch := make(chan string)
	go func() {
		fmt.Println("Hello world...inside goroutine...")
	}()

	fmt.Println("Exiting now...")
}

func main() {
	ch := make(chan string)
	go func() {
		fmt.Println("Hello world...inside goroutine...")
		ch <- "Hey Channel"
	}()

	<-ch // This is blocking call. <-ch will not finish until channel has some value to emit out.
	// Thus it will wait to till channel gets some value and then it emit out.
	// <- ch will forever wait if no one put anything in the channel.
	fmt.Println("Exiting now...")
}
