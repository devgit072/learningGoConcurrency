package main

import (
	"fmt"
	"strings"
	"sync"
)

// Below is example of callback.
// Callback is nothing but when function is passed as parameter, then it is called callback.
// Below example is sync callback.
func main1() {
	printerFunc := func(msg string) {
		fmt.Println("I'm printer function, I will just print anything given...")
		fmt.Println("Printing: ", msg)
	}
	toUpperCase("hellWorld_cacvacac  kkkk KKKKK", printerFunc)
}

func toUpperCase(word string, printerFunc func(string)) {
	uperCaseWord := strings.ToUpper(word)
	printerFunc(uperCaseWord)
}
//==================================================================================================

// Lets see the example of Async callback which can be achived using go routine.
func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	printerFunc := func(msg string) {
		fmt.Println("Printing: ", msg)
		waitGroup.Done()
	}
	toUpperAsync("Hello World", printerFunc)
	fmt.Println("Waiting for callback......")
	waitGroup.Wait()
	fmt.Println("exiting main...")
}
func toUpperAsync(word string, printerFunc func(string)) {
	upperCaseWord := strings.ToUpper(word)
	go func() {
		printerFunc(upperCaseWord)
	}()
}





