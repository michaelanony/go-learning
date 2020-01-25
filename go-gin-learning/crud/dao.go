package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDb()(err error)  {
	addr :="michael:cccbbb@tcp(192.168.11.31:30001)/testDb"
	db, err = sqlx.Connect("mysql", addr)
	if err!=nil{
		return err
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return
}

func QueryAllBook()(bookList []*Book,err error)  {
	sqlStr:="select * from book_test"
	err = db.Select(&bookList,sqlStr)
	if err!=nil{
		fmt.Println("Query failed")
		return
	}
	return
}

func InsertBook(title string,price int64) (err error) {
	sqlStr :="insert into book_test(title,price) values(?,?)"
	_, err = db.Exec(sqlStr, title, price)
	if err!=nil{
		fmt.Println("insert failed")
		return
	}
	return
}

func DeleteBook(book string) (err error){
	sqlStr :="delete from book_test where title =?"
	_, err = db.Exec(sqlStr, book)
	if err!=nil{
		fmt.Println("delete failed")
		return
	}
	return
}