package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		//time.Sleep(time.Second)
		ch <- 2
		close(ch)// closing channel means channel cant take value anymore.
	}()

	for v := range ch {
		fmt.Println("Here is chan val", v)
	}

}


