package bot

import (
    //"fmt"
    "log"
    "net/http"
    "github.com/PuerkitoBio/goquery"
    //"strconv"
    //"github.com/bwmarrin/discordgo"
)
var(
    a []string
    b []byte
)
func Crawling(str string,num string) ([]string){
    a = []string{}
    b = []byte{}
	response, err := http.Get("https://danbooru.donmai.us/posts?page=" + num + "&tags=" + str)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }
    document.Find("article").Each(func(i int,e *goquery.Selection){
        href,exist := e.Attr("data-file-url")
        if exist{
            //b = []byte(href)[6:]
            a = append(a,/*string(b)*/href + "\n")
        }
    })
    return a
}
