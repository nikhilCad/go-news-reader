package main

import (
	"fmt"
	"os"
	
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/mmcdole/gofeed"
)

func main() {

	urls := []string{
		"http://feeds.twit.tv/twit.xml",
		"https://rss.nytimes.com/services/xml/rss/nyt/World.xml",
	}

	var fp = gofeed.NewParser()

	items := []list.Item{}
	
	for _, url := range urls {

		var feedParsed, _ = fp.ParseURL(url)

		fmt.Println(feedParsed.Title)

		for _, parsedItem := range feedParsed.Items {

			items = append(items, item{
				title: parsedItem.Title,
				desc: parsedItem.Description,
				url:   parsedItem.Link,
			})

		}
	}

	m := model{list: list.New(items, newCustomDelegate(), 0, 0)}
	m.list.Title = "News"
	m.viewport = viewport.New(0, 0)  // Initial dimensions, will be set later

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	
}