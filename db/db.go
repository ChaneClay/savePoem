package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var(
	db *sql.DB
	err error
)

const(
	MaxConns int = 20
	MixConns int = 2
)

func init()  {
	//db, err := sql.Open("mysql", "go:go@tcp(127.0.0.1:336)/poem?charset=utf-8&parseTime=true")
	db, _ = sql.Open("mysql", "root:111111@tcp(127.0.0.1:3306)/go?charset=utf8")
	if err!=nil{
		panic(err)
	}
	db.SetMaxIdleConns(MaxConns)
	db.SetMaxOpenConns(MixConns)
	err = db.Ping()
	if err != nil{
		panic(err)
	}


	
}