package db

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
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


var file = xlsx.NewFile()
var sheet, _ = file.AddSheet("Sheet1")
var row = sheet.AddRow()			//第一行
var num=0
var cell = row.AddCell()

func (p *Poem) SaveToXlsx()bool{
	var Title = p.Title
	var Author = p.Author
	var Dynasty = p.Dynasty
	var Content = p.Content
	var path = "F:\\testData\\hh\\"

	if num == 0{
		row.SetHeightCM(1)
		cell.Value = "Title"
		cell = row.AddCell()
		cell.Value = "Author"
		cell = row.AddCell()
		cell.Value = "Dynasty"
		cell = row.AddCell()
		cell.Value = "Content"
		num ++
	}

	row = sheet.AddRow()			//第二行
	//row.SetHeightCM(0.8)

	cell = row.AddCell()
	cell.Value = Title
	cell = row.AddCell()
	cell.Value = Author
	cell = row.AddCell()
	cell.Value = Dynasty
	cell = row.AddCell()
	cell.Value = Content

	err = file.Save(path+"poemData.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
		return false
	}
	return true

}


func (p *Poem) Save(){
	json.Marshal(p)
	//res := p.Insert()
	res := p.SaveToXlsx()
	if !res {
		fmt.Println("插入失败。。。")
	}

	//fmt.Println(string(data))
}
