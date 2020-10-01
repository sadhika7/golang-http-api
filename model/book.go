package model

import "log"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
