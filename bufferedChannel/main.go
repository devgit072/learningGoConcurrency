package main

import (
	"fmt"
	"net/http"
	"sync"
)

// This are worker who will get the fetch the url and print response and say Done().
func webGetWorker(in <- chan string, wg *sync.WaitGroup) {
	for {
		url := <- in
		res, err := http.Get(url)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Get %s, status: %d", url, res.StatusCode)
		}
		wg.Done()
	}
}

func main() {

	bufferedCh := make(chan string, 1000)
	var wg sync.WaitGroup

	numOfWorker := 4
	for i:= 0 ; i< numOfWorker ;i++ {
		go webGetWorker(bufferedCh, &wg)
	}

	urls := []string {"http://google.com", "http://rubrik.com", "http://cohesity.com", "http://apple.com", "http://facebook.com", "http://udemy.com"}

	for i:=0 ;i< 100 ;i++ {
		for _,url := range urls {
			wg.Add(1)
			bufferedCh <- url // this is non-blocking since it is buffered channel. It will keep putting in channel untill it reaches full capacity.

		}
	}
	wg.Wait()
}
