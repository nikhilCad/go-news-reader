package main

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

// struct fields need to be uppercase, json field lowercase is just preference
type item struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image string `json:"image"`
}

type feed struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Link string `json:"link"`
	Items []item `json:"items"`
}

var feeds = populateFeeds()

func populateFeeds() []feed {

	urls := []string{
		"http://feeds.twit.tv/twit.xml",
		"https://rss.nytimes.com/services/xml/rss/nyt/World.xml",
	}

	var feeds []feed

	var fp = gofeed.NewParser()

	
	for _, url := range urls {

		var feedParsed, _ = fp.ParseURL(url)

		newFeed := feed{
			Title:       feedParsed.Title,
			Description: feedParsed.Description,
			Link:        feedParsed.Link,
		}

		for _, parsedItem := range feedParsed.Items {
			
			imgUrl := ""
			if parsedItem.Image != nil && parsedItem.Image.URL != "" {
				imgUrl = parsedItem.Image.URL
			}

			newItem := item{
				Title:       parsedItem.Title,
				Link:        parsedItem.Link,
				Description: parsedItem.Description,
				Image: imgUrl,
			}

			newFeed.Items = append(newFeed.Items, newItem)
		}

		feeds = append(feeds, newFeed)
	}

	return feeds

}

func getFeeds(context *gin.Context){
	context.IndentedJSON(http.StatusOK, feeds)
}

func createFeed(c *gin.Context) {
	var newFeed feed
	if err := c.BindJSON(&newFeed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	feeds = append(feeds, newFeed)
	c.IndentedJSON(http.StatusCreated, newFeed)
}

func main() {
	router := gin.Default()
	router.GET("/feeds", getFeeds)
	router.POST("/feeds", createFeed)
	router.Run("localhost:8090")
}
