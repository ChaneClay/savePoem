package handle

import (
	"PoemGet/gofish"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
)

var baseUrl = "https://so.gushiwen.org/"

type Handle interface {
	Work(body io.Reader, url string)
}

type AuthorHandle struct {}

func (h *AuthorHandle) Worker(body io.Reader, url string)  {

	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil{
		fmt.Println("doc.err.",err)
	}
	//fmt.Println(doc.Text())
	doc.Find(".sons").Find(".cont").Find("a").Each(func(i int, s *goquery.Selection) {
		author := s.Text()
		fmt.Println("正在下载： ", author)
		//fmt.Printf("%d author=%s", i, author)
		link, _ := s.Attr("href")
		//fmt.Printf(" link=%s\n",link)

		h := PoemHomeHandle{}
		fish := gofish.NewGoFish()
		request, err := gofish.NewRequest("GET", baseUrl+link, gofish.UserAgent, &h, nil)

		if err != nil{
			fmt.Println(err)
			return
		}
		fish.Request = request
		fish.Visit()
	})
}

type PoemHomeHandle struct {}

func (h *PoemHomeHandle) Worker(body io.Reader, url string) {
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil{
		fmt.Println("doc.err.",err)
	}

	doc.Find(".sonspic").Find(".cont").Find("p").Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		fmt.Println("作品主页：", baseUrl+link)


		h := PoemInfoHandle{}
		fish := gofish.NewGoFish()
		request, err := gofish.NewRequest("GET", baseUrl+link, gofish.UserAgent, &h, nil)

		if err != nil{
			fmt.Println(err)
			return
		}
		fish.Request = request
		fish.Visit()


	})
}









