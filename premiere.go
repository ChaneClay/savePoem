package main

import (
	"PoemGet/gofish"
	"PoemGet/handle"
	"fmt"
)


func main()  {
	authors := "https://so.gushiwen.org/authors/"

	h := handle.AuthorHandle{}
	fish := gofish.NewGoFish()
	request, err := gofish.NewRequest("GET", authors, gofish.UserAgent, &h, nil)

	if err != nil{
		fmt.Println(err)
		return
	}
	fish.Request = request
	fish.Visit()
}












