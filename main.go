package main

import (
	"fmt"
	// "net/http"
	// "github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

func main() {

	urls := []string{
		"http://feeds.twit.tv/twit.xml",
		"https://rss.nytimes.com/services/xml/rss/nyt/World.xml",
	}

	var fp = gofeed.NewParser()

	for _, url := range urls {

		var feedParsed, _ = fp.ParseURL(url)

		fmt.Println(feedParsed.Title)
	}

	
}