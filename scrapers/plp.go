package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const Letters = "abcdefghijklmnopqrstuvwxyz"
const BaseUrl = "http://www.portaldalinguaportuguesa.org"

func getWords(letter string) map[string]string {

	rootCollector := colly.NewCollector(
		colly.Async(true),
	)

	rootCollector.SetRequestTimeout(time.Second * 10000)

	var wordsMap = make(map[string]string)
	var locker sync.Mutex

	rootCollector.OnError(func(res *colly.Response, err error) {
		err = res.Request.Retry()
		if err != nil {
			fmt.Println("Error retrying request", err)
			return
		}
	})

	rootCollector.OnHTML("table#rollovertable > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			w := strings.Split(el.ChildText("td:nth-child(1)"), " ")[0]
			w = strings.Trim(w, " \t\n\r")
			html, _ := el.DOM.Find("td:nth-child(2)").Html()
			html = strings.ReplaceAll(html, "<b>", "")
			html = strings.ReplaceAll(html, "</b>", "")
			html = strings.ReplaceAll(html, "</u>", "")
			html = strings.ReplaceAll(html, "<u>", "'")
			html = strings.ReplaceAll(html, "<span class=\"nobreaksyllable\">", "")
			html = strings.ReplaceAll(html, "<span style=\"color: #aaaaaa;\">", "")
			html = strings.ReplaceAll(html, "</span>", "")
			html = strings.ReplaceAll(html, "·", "-")
			html = strings.Trim(html, " \t\n\r")

			if strings.ContainsAny(html, "><") {
				panic(html)
			}
			if len(w) == 0 || len(html) == 0 {
				return
			}

			locker.Lock()
			if _, ok := wordsMap[w]; !ok {
				wordsMap[w] = html
			}
			locker.Unlock()
		})
	})

	rootCollector.OnHTML("#maintext > p", func(e *colly.HTMLElement) {
		parts := strings.Split(e.Text, " ")
		if len(parts) < 5 {
			return
		}
		checkValue, _ := strconv.Atoi(parts[2])

		if checkValue != 999999999 {
			return
		}

		nWords, _ := strconv.Atoi(strings.Split(e.Text, " ")[4])

		for p := 0; p < nWords; p += 100 {
			rootCollector.Visit(fmt.Sprintf("%v/index.php?action=syllables&act=list&letter=%v&start=%v", BaseUrl, letter, p))

			if nWords%100 == 0 {
				rootCollector.Wait()
				time.Sleep(time.Millisecond * 200)
			}
		}

	})

	rootCollector.Visit(fmt.Sprintf("%v/index.php?action=syllables&act=list&letter=%v&start=999999999", BaseUrl, letter))
	rootCollector.Wait()

	return wordsMap
}

func main() {
	f, err := os.Create("words.csv")
	defer f.Close()

	if err != nil {
		panic(err)
	}

	for _, l := range Letters {
		words := getWords(string(l))
		println(string(l), len(words))
		for k, v := range words {
			_, err := f.WriteString(fmt.Sprintf("%v,%v\n", k, v))
			if err != nil {
				panic(err)
			}
		}
	}

	println("Done")
}
