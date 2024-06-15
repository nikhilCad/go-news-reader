package main

import (
	"log"
	"time"

	readability "github.com/go-shiori/go-readability"
)

func ParseUrl(url string) string {
	article, err := readability.FromURL(url, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to parse %s, %v\n", url, err)
		return " Failed to parse"
	}

	return article.TextContent


}