package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //db连接池
type user struct {
	id           int
	city         string
	abbreviation string
}

func initDB() (err error) {
	dsn := ""
	db, err := sql.Open("mysql", dsn) //不检查用户密码
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
func query() {

}
func insert() {

}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init failed%v", err)
	}
	fmt.Println("success")
	//1、写查询单个记录的sql语句
	sqlStr := `select id,city,abbreviation from city_abbreviation where id = 4549`
	//2、执行
	ret := db.QueryRow(sqlStr)
	//3、拿到结果
	var u1 user
	ret.Scan(&u1.id, &u1.city, &u1.abbreviation)
	fmt.Printf("%v", u1)
}
