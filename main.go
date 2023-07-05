package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func WebScrap() {
	resp, err := http.Get("https://ru.hexlet.io/blog/posts/participate-in-open-source?utm_source=hexlet&utm_medium=blog&utm_campaign=hexlet-blog&utm_content=open-sourse-ne-strashen&utm_term=post_261022")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Status code error %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
	fmt.Println(doc)
}
func main() {
	WebScrap()

}
