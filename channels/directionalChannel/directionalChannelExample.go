package main

import (
	"fmt"
)

// channel can be stated as recieving or emitting channel.

func main1() {
	channel := make(chan string, 1)
	go func(ch chan <-string) { // Here we have stated that ch is recieving channel
		ch <- "Hello World"
		// val := <-ch will retruns error because in parameter itself we have mentioned that ch is recieving channel.
	}(channel)

	message := <-channel
	fmt.Println(message)
}

//==============================================================
func main() {
	channel := make(chan string, 1)
	go func(ch chan <-string) { // Here we have stated that ch is recieving channel
		ch <- "Hello World"
		// val := <-ch will retruns error because in parameter itself we have mentioned that ch is recieving channel.
	}(channel)

	message := channelHavingEmitOnly(channel)
	fmt.Println(message)
}


func channelHavingEmitOnly(ch <- chan string) string { // this channel will only emit and cant recieve any thing
	val := <-ch
	return val
}

