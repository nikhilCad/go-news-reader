package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabae() {
	database, _ := sql.Open("sqlite3", "./feeds.db")
	
	var sqlCreate = "CREATE TABLE IF NOT EXISTS feeds (url VARCHAR PRIMARY KEY)"
	statement,_ := database.Prepare(sqlCreate)
	statement.Exec()
	
	database.Close()
}

func UpdateDatabase() {
	database, _ := sql.Open("sqlite3", "./feeds.db")
	
	statement, _ := database.Prepare("INSERT INTO feeds (url) VALUES (?)")
	statement.Exec("https://timesofindia.indiatimes.com/rssfeedstopstories.cms")
	
	database.Close()
}

func ReadDatabase() []string {
	database, _ := sql.Open("sqlite3", "./feeds.db")
	rows, _ := database.Query("SELECT url FROM feeds")

	var url string
	var urls []string
	for rows.Next() {
		rows.Scan(&url)
		urls = append(urls, url)
	}

	database.Close()

	return urls
}