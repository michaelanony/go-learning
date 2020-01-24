package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //db连接池
type user struct {
	id           int
	city         sql.NullString
	abbreviation sql.NullString
}

func initDB() (err error) {
	dsn := "michael:cccbbb@tcp(192.168.11.31:30001)/spider_12306"
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
func queryOne(id int) {
	var u1 user
	//1、写查询单个记录的sql语句.查询单个数据
	sqlStr := `select id,city,abbreviation from city_abbreviation where id = ?;`
	//2、执行
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.city, &u1.abbreviation)
	//3、拿到结果
	fmt.Printf("%v", u1)
}
func queryMore(n int) {
	sqlStr := `select id,city,abbreviation from city_abbreviation where id >?;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.city, &u1.abbreviation)
		if err != nil {
			fmt.Printf("%v", err)
		}
		fmt.Printf("%v", u1)
	}
}

//预处理插入多条数据
func insert() {
	sqlStr := `insert into user(name,age) values(?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("failed")
	}
	defer stmt.Close()
	var m = map[string]int{
		"lq": 30,
		"wq": 32,
		"ts": 21,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

//mysql事务处理
func transaciton() {
	//1.开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("failed")
		return
	}
	sqlStr1 := `update user set age=age-2 where id =1`
	sqlStr2 := `update user set age=age+2 where id =2`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init failed%v", err)
	}
	fmt.Println("success")
	//queryOne(4555)
	queryMore(4000)

	//查询多条数据
}
