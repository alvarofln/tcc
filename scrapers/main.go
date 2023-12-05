package main

import (
	"github.com/gocolly/colly"
	html2 "html"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getArtistsLinks() []string {
	rootCollector := colly.NewCollector()
	links := make([]string, 0, 1024)
	rootCollector.OnHTML(".moreNamesContainer > ul > li > a", func(e *colly.HTMLElement) {
		links = append(links, e.Attr("href"))
	})
	rootCollector.Visit("https://www.vagalume.com.br/browse/style/sertanejo.html")
	rootCollector.Wait()

	return links
}

func dumpLyric(link string) {
	rootCollector := colly.NewCollector(
		colly.Async(true),
	)

	rootCollector.SetRequestTimeout(time.Second * 5)

	rootCollector.OnError(func(res *colly.Response, err error) {
		if res.StatusCode != http.StatusNotFound {
			res.Request.Retry()
		}
	})

	rootCollector.OnHTML("#body", func(e *colly.HTMLElement) {
		requestPath := e.Request.URL.Path

		lang, langExists := e.DOM.Find(".lang").Attr("class")

		if langExists && !strings.Contains(lang, "-bra") {
			println(requestPath + " non pt lyric detected, skipping...")
		}

		html, _ := e.DOM.Find("#lyrics").Html()
		content := html2.UnescapeString(strings.ReplaceAll(html, "<br/>", "\n"))
		lyricPath := "./letras" + strings.ReplaceAll(requestPath, ".html", ".txt")
		f, err := os.Create(lyricPath)

		if err != nil {
			log.Fatal(err)
		}
		_, err = f.WriteString(content + "\n")

		if err != nil {
			log.Fatal(err)
		}

		f.Sync()
	})

	rootCollector.Visit(link)
	print(link + " crawling...")
	rootCollector.Wait()
	println("done")
}

func dumpArtistLyrics(artistLink string, topOnly bool) {
	path := "./letras" + artistLink
	os.MkdirAll(path, os.ModePerm)

	rootCollector := colly.NewCollector(
		colly.Async(true),
	)

	rootCollector.SetRequestTimeout(time.Second * 5)

	rootCollector.OnError(func(res *colly.Response, err error) {
		if res.StatusCode != http.StatusNotFound {
			res.Request.Retry()
		}
	})

	tagId := "#alfabetMusicList"

	if topOnly {
		tagId = "#topMusicList"
	}

	rootCollector.OnHTML(tagId, func(e *colly.HTMLElement) {
		for _, link := range e.ChildAttrs(".nameMusic", "href") {
			rootCollector.Visit("https://www.vagalume.com.br" + link)
		}
	})

	rootCollector.OnHTML("#body", func(e *colly.HTMLElement) {
		requestPath := e.Request.URL.Path

		lang, langExists := e.DOM.Find(".lang").Attr("class")

		if langExists && !strings.Contains(lang, "-bra") {
			println(requestPath + " non pt lyric detected, skipping...")
			return
		}

		html, _ := e.DOM.Find("#lyrics").Html()

		if html == "" {
			return
		}

		content := html2.UnescapeString(strings.ReplaceAll(html, "<br/>", "\n"))
		lyricPath := "./letras" + strings.ReplaceAll(requestPath, ".html", ".txt")
		f, err := os.Create(lyricPath)

		if err != nil {
			log.Fatal(err)
		}
		_, err = f.WriteString(content + "\n")

		if err != nil {
			log.Fatal(err)
		}

		f.Sync()
	})

	rootCollector.Visit("https://www.vagalume.com.br" + artistLink)
	print(artistLink + " crawling...")
	rootCollector.Wait()
	println("done")

}

// function read all lines from txt files

func joinAllLyrics() {
	f, _ := os.OpenFile("./out.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	filepath.Walk("./letras", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			bytes, _ := os.ReadFile(path)
			content := string(bytes)
			f.WriteString(content)
			f.Sync()
		}
		return nil
	})

}

func main() {
	for _, artistLink := range getArtistsLinks() {
		dumpArtistLyrics(artistLink, false)
	}
	joinAllLyrics()
	println("done")
}
