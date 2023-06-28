package main

import (
	"backend-course-go/section2"
	"log"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	logger := log.Default()
	client := section2.NewClient(section2.BaseUrl, 3, logger)
	ads, _ := client.GetAdByCate()
	logger.Println("Total %d", ads.Total)
}
