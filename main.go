package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/muesli/termenv"
)

// The submission struct
type Submission struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

// Page number
var pageNum int = 1

// List of links and titles (For feeding the TUI part)
var submissions []Submission = []Submission{}

// Helper functions
func checkStatusCode(msg string, res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("%s | %d %s", msg, res.StatusCode, res.Status)
	}
}

// Scrape the data at page number 'page'
func Scrape(page int) []Submission {
	// Add some text and visual effects before we actually scrape
	termenv.ClearScreen()
	termenv.AltScreen()

	fmt.Println("Waiting...")

	// Scrape the news at the page number 'pageNum'
	res, err := http.Get(spf("https://news.ycombinator.com/news?p=%d", pageNum))
	check_err(err)
	defer res.Body.Close()
	checkStatusCode("Status code error", res)

	// Create a document for scraping
	doc, err := goquery.NewDocumentFromReader(res.Body)
	check_err(err)

	// Scrape the submissions
	doc.Find("tr .athing").Each(func(i int, s *goquery.Selection) {
		// Take the ID of the submission (very important for scraping the submissions data)
		id, _ := s.Attr("id")

		// Scrape the submissions data using the ID from Hacker New's Firebase database
		jsonRes, err := http.Get(spf("https://hacker-news.firebaseio.com/v0/item/%s.json?print=pretty", id))
		check_err(err)
		defer jsonRes.Body.Close()
		checkStatusCode(spf("Cannot scrape data of ID %s", id), jsonRes)

		doc, err := goquery.NewDocumentFromReader(jsonRes.Body)
		check_err(err)
		jsonData := doc.Children().Text()

		var submission Submission

		// Unmarshal the JSON data to Submission
		err = json.Unmarshal([]byte(jsonData), &submission)
		check_err(err)

		// In case that the submission doesn't have an URL (like with a local post and not a link)
		// We will just give the post link instead
		if submission.URL == "" {
			submission.URL = spf("https://news.ycombinator.com/item?id=%s", id)
		}

		submissions = append(submissions, submission)
	})

	// Exit the alt screen
	termenv.ExitAltScreen()

	return submissions
}

func main() {
	submissions = Scrape(pageNum)
	tui()
}
