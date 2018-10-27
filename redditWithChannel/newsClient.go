package main

import (
	"fmt"
	"github.com/caser/gophernews"
	"github.com/jzelinskie/geddit"
	"os"
	"sync"
)

var redditSession *geddit.LoginSession
var hackerNewsClient *gophernews.Client

func init() {
	hackerNewsClient = gophernews.NewClient()
	var err error
	redditSession, err = geddit.NewLoginSession("a_monk_coder", "gopherofbihar", "gdAgent v0")

	if err != nil {
		panic(err)
	}
}

type Story struct {
	title string
	url string
	source string
	author string
}


func getHnStoryDetails(id int, c chan <- Story, wg *sync.WaitGroup) {
	defer wg.Done()

	story, err := hackerNewsClient.GetStory(id)
	if err != nil {
		return
	}
	newStory := Story{
		title: story.Title,
		url:story.URL,
		author:story.By,
		source: "HackerNews",
	}

	c <- newStory
}

func newHnStories(c chan <- Story) {
	defer close(c)
	var wg sync.WaitGroup
	changes, err := hackerNewsClient.GetChanges()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, id := range changes.Items {
		wg.Add(1)
		go getHnStoryDetails(id, c, &wg)
	}
	wg.Wait()
}

func newRedditStories(c chan <- Story) []Story {
	defer close(c)

	var stories []Story
	sort := geddit.PopularitySort(geddit.NewSubmissions)
	var listingOptions geddit.ListingOptions
	submissions, err := redditSession.SubredditSubmissions("golang", sort, listingOptions)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, submission := range submissions {
		newStory := Story{
			title: submission.Title,
			url : submission.URL,
			author:submission.Author,
			source:"Reddit /r/golang",
		}
		c <- newStory
	}

	return stories
}

func outPutToConsole(c <- chan Story) {
	for {
		s := <- c
		str := fmt.Sprintf("%s", s)
		fmt.Println(str)
	}
}

func outPutToFile(c <- chan Story, file *os.File) {
	for {
		s := <- c
		fmt.Fprintf(file, "%s\n", s)
	}
}


func main() {

	hnChan := make(chan Story, 8)
	ged := make(chan Story, 8)
	toFile := make(chan Story, 8)
	toConsole := make(chan Story, 8)
	go newRedditStories(ged)
	go newHnStories(hnChan)

	file, err := os.Create("stories.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go outPutToConsole(toConsole)
	go outPutToFile(toFile, file)

	defer file.Close()

	hnOpen := true
	redditOpen := true

	for hnOpen || redditOpen {
		select {
		case story, open := <- hnChan:
			if open {
				toConsole <- story
				toFile <- story
			} else {
				hnOpen = false
			}
		case story, open := <-ged:
			if open {
				toConsole <- story
				toFile <- story
			} else {
				redditOpen = false
			}
		}
	}
}
