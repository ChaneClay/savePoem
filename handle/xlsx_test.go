package handle

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"testing"
)

func TestXlsx(t *testing.T) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row, row1, row2 *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	row.SetHeightCM(1)

	cell = row.AddCell()
	cell.Value = "姓名"
	cell = row.AddCell()
	cell.Value = "年龄"

	row1 = sheet.AddRow()
	row1.SetHeightCM(1)
	cell = row1.AddCell()
	cell.Value = "狗子"
	cell = row1.AddCell()
	cell.Value = "18"

	row2 = sheet.AddRow()
	row2.SetHeightCM(1)
	cell = row2.AddCell()
	cell.Value = "蛋子"
	cell = row2.AddCell()
	cell.Value = "28"

	err = file.Save("test_write.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}


func TestXlsx2(t *testing.T) {

	var Title = "将进酒"
	var Author = "李白李白李白李白李白"
	var Dynasty = "唐代"
	var Content = `君不见，黄河之水天上来，奔流到海不复回。
	君不见，高堂明镜悲白发，朝如青丝暮成雪。
	人生得意须尽欢，莫使金樽空对月。
	天生我材必有用，千金散尽还复来。
	烹羊宰牛且为乐，会须一饮三百杯。
	岑夫子，丹丘生，将进酒，杯莫停。
	与君歌一曲，请君为我倾耳听。(倾耳听 一作：侧耳听)
	钟鼓馔玉不足贵，但愿长醉不复醒。(不足贵 一作：何足贵；不复醒 一作：不愿醒)
	古来圣贤皆寂寞，惟有饮者留其名。(古来 一作：自古；惟 通：唯)
	陈王昔时宴平乐，斗酒十千恣欢谑。
	主人何为言少钱，径须沽取对君酌。
	五花马，千金裘，呼儿将出换美酒，与尔同销万古愁。`

	var path = "F:\\testData\\hh\\"
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row := sheet.AddRow()			//第一行
	row.SetHeightCM(1)
	cell := row.AddCell()
	cell.Value = "Title"
	cell = row.AddCell()
	cell.Value = "Author"
	cell = row.AddCell()
	cell.Value = "Dynasty"
	cell = row.AddCell()
	cell.Value = "Content"

	for i := 0; i<5; i++{
		row = sheet.AddRow()			//第二行
		row.SetHeightCM(1)
		cell = row.AddCell()
		cell.Value = Title
		cell = row.AddCell()
		cell.Value = Author
		cell = row.AddCell()
		cell.Value = Dynasty
		cell = row.AddCell()
		cell.Value = Content
	}



	err = file.Save(path+"test_write.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
