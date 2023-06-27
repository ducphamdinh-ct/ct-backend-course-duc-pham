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
	classes, err := section2.ReadJson()
	err = section2.WriteFile(classes)
	log.Fatal(err)
}
