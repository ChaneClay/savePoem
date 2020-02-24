package db

import (
	"encoding/json"
)

type Poem struct {
	Id int
	Title string
	Author string
	Dynasty string
	Content string
}

func (p *Poem) Insert() bool{
	stmtInsert, err := db.Prepare("insert into poem(title, author, dynasty, content)values(?,?,?,?)")
	if err != nil{
		return false
	}
	_, err = stmtInsert.Exec(&p.Title, &p.Author, &p.Dynasty, &p.Content)
	if err != nil{
		return false
	}
	return true

}

func (p *Poem) Save(){
	json.Marshal(p)
	p.Insert()
	//fmt.Println(string(data))
}
