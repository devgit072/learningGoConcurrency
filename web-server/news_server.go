package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Story struct {
	title string
	url string
	source string
	author string
}

var stories []Story

func init() {

	stories = append(stories,
		Story{"Go language Story", "http://example.com", "PacktViewer", "LearningGo"},
		Story{"Go performance Story", "http://example.com", "PacktViewer", "LearningGo"},
		Story{"Rust language Story", "http://example.com", "PacktViewer", "LearningGo"},
		Story{"Python language Story", "http://example.com", "PacktViewer", "LearningGo"},
		Story{"C++ language Story", "http://example.com", "PacktViewer", "LearningGo"},
		)
}


func searchStory(query string) []Story {
	var foundStory []Story
	for _, story := range stories {
		if strings.Contains(strings.ToUpper(story.title), strings.ToUpper(query)) {
			foundStory = append(foundStory, story)
		}
	}

	return foundStory
}

func search(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	if query == "" {
		http.Error(w, "Search parameter is empty", http.StatusNotAcceptable)
		return
	}

	w.Write([]byte("<html><body>"))
	s := searchStory(query)

	if len(s) == 0 {
		w.Write([]byte(fmt.Sprintf("No result for query: %s . \n <br>" , query)))
	} else {
		for _, story := range  s {
			w.Write([]byte(fmt.Sprintf("<a href='%s'>%s</a><br>by %s on %s <br><br>", story.url, story.title,
				story.author, story.source)))
		}
	}
	w.Write([]byte("<a href='../'>Back</a>"))
	w.Write([]byte("</body></html>"))
}

func topTen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><body>"))
	form := "<form action='search' method='get'> Search: <input type ='text' name='q'> <input type='submit'></form>"
	w.Write([]byte(form))

	for i := len(stories) - 1;i >= 0 && len(stories)-i<10;i-- {
		story := stories[i]
		w.Write([]byte(fmt.Sprintf("<a href='%s'>%s</a><br>by %s on %s <br><br>", story.url, story.title,
			story.author, story.source)))
	}

	w.Write([]byte("</body></html>"))
}

func main() {
	http.HandleFunc("/", topTen)
	http.HandleFunc("/search", search)

	if err := http.ListenAndServe(":8080", nil) ; err != nil {
		panic(err)
	}
}