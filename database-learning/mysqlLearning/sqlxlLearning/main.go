package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type staticNum struct {
	Sproject int
	Id       int
	Szhuanli int
	Smanyi   int
}

func initDB() (err error) {
	dsn := "michael:cccbbb@tcp(192.168.11.31:30001)/home"
	db, err := sqlx.Connect("mysql", dsn) //不检查用户密码
	if err != nil {
		fmt.Printf("faild,err:%v", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("faild,err:%v", err)
		return
	}
	fmt.Println("Success")
	return
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	sqlStr := `select * from lywebback_static_num where id =1`
	var u staticNum
	db.Get(&u, sqlStr)
	fmt.Printf("%v", u)
}
