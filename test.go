package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)
func processElement(index int, element *goquery.Selection) {
    href, exists := element.Attr("id")
    if exists {
        
    }
}

func main() {
    response, err := http.Get("https://danbooru.donmai.us/posts?tags=feet")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }
    return document.Find("article").Each(processElement)
}