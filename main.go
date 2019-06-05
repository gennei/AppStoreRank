package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://itunes.apple.com/jp/app/id359066582?mt=8"

	doc, _ := goquery.NewDocument(url)
	fmt.Println(fixTitle(doc.Find(".inline-list__item").First().Text()))
}

func fixTitle(title string) string {
	return strings.TrimSpace(title)
}
