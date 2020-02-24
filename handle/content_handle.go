package handle

import (
	"PoemGet/db"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"strings"
)

var poemHome = "https://so.gushiwen.org/authors/authorvsw_b90660e3e492A1.aspx"

func getUrls(url string, size int) []string{
	urls := make([]string, 0)
	urlTpl := strings.Replace(url, "A1.aspx", "A%d.aspx", 1)

	for i:=1; i<=size; i++{
		urls = append(urls, fmt.Sprintf(urlTpl, i))
		//fmt.Println(urls)
		//fmt.Println(fmt.Sprintf(urlTpl, i))
	}
	//fmt.Printf("len: %d, cap: %d\n", len(urls), cap(urls))
	return urls
}


type PoemInfoHandle struct {}

func (h *PoemInfoHandle) Worker(body io.Reader, url string) {
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil{
		fmt.Println("doc.err.",err)
	}

	doc.Find(".cont").Each(func(i int, s *goquery.Selection) {
		author := ""
		dynasty := ""
		content := ""
		title := ""

		title = strings.TrimSpace(s.Find("p").Find("a").Find("b").Text())

		authorAndDynasty := strings.TrimSpace(s.Find(".source").Text())

		authorAndDynastySlice := strings.Split(authorAndDynasty, "：")

		if len(authorAndDynastySlice) == 2{
			dynasty = authorAndDynastySlice[0]
			author = authorAndDynastySlice[1]
		}
		s.Find(".contson").Each(func(i int, s *goquery.Selection) {
			content = strings.TrimSpace(s.Text())
		})

		//fmt.Printf("author: %s,dynasty: %s,content: %s,title: %s", author, dynasty, content, title)

		if author != "" && dynasty != "" && title != "" && content != "" {
			p := db.Poem{}
			p.Author = author
			p.Title = title
			p.Content = content
			p.Dynasty = dynasty
			p.Save()
		}
	})
}

