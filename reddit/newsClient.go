package main

import (
	"fmt"
	"github.com/caser/gophernews"
	"github.com/jzelinskie/geddit"
	"os"
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

func newHnStories() []Story {
	var stories []Story
	changes, err := hackerNewsClient.GetChanges()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, id := range changes.Items {
		story, err := hackerNewsClient.GetStory(id)
		if err != nil {
			continue
		}
		newStory := Story{
			title: story.Title,
			url:story.URL,
			author:story.By,
			source: "HackerNews",
		}

		stories = append(stories, newStory)
	}

	return stories
}

func newRedditStories() []Story {
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
		stories = append(stories, newStory)
	}

	return stories
}

func main() {

	hnStories := newHnStories()
	redditStories := newRedditStories()

	var stories []Story

	if hnStories != nil {
		stories = append(stories, hnStories...)
	}

	if redditStories != nil {
		stories = append(stories, redditStories...)
	}

	file, err := os.Create("stories.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	for _, s := range stories {
		str := fmt.Sprintf("%s\n", s)
		fmt.Fprintf(file, str)
	}

	for _, s := range stories {
		fmt.Println(file, s)
	}
}
