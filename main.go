package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func WebScrap() {
	var adresss = "https://www.study.ru/courses/elementary/osnovy"
	urlParse, err := url.Parse(adresss)
	if err != nil {
		return
	}
	for len(adresss) > 0 {

		resp, err := http.Get(urlParse.String())
		if err != nil {
			log.Fatalf("Failed to get the webpage: %v", err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Fatalf("Status code error %d %s", resp.StatusCode, resp.Status)
		}
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Print(err)
		}
		doc.Find("div.table-wrapper").Each(func(i int, s *goquery.Selection) {
			title := s.Text()
			fmt.Printf("Review %d: %s\n", i, title)
		})
		doc.Find("a.button-fill").Each(func(i int, s *goquery.Selection) {
			adress, ok := s.Attr("href")
			if !ok {
				return
			}
			relativeUrl, err := url.Parse(adress)
			if err != nil {
				fmt.Printf("failed parese URL %v", err)
			}
			urlParse = urlParse.ResolveReference(relativeUrl)
			fmt.Println(urlParse)
			adress = urlParse.String()
		})
	}
}

func main() {
	WebScrap()
}
