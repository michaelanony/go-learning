package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	var err error
	JsonParse := NewJsonStruct()
	v := MainConfig{}
	JsonParse.Load("./config.json", &v)
	DB, err = sqlx.Open(`mysql`,
		v.User+`:`+v.Pwd+`@tcp(`+v.Host+`:`+v.Port+`)/home?charset=utf8&parseTime=true`)
	if err != nil {
		panic("Connection Error!")
	}
	if err = DB.Ping(); err != nil {
		panic("Run Error")
	}
}

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Num      string `json:"num" form:"num"`
	Password string `json:"password" form:"password"`
	Logo     string `json:"logo" form:"logo"`
	Age      int    `json:"age" form:"age"`
}

type LyStaticNum struct {
	Id       int
	Smanyi   int
	Sproject int
	Szhuanli int	`json:"szhuanli" form:"szhuanli"`
}

func lyStaticNumAll() ([]LyStaticNum, error) {
	mods := make([]LyStaticNum, 0)
	err := DB.Select(&mods, "SELECT * FROM lywebback_static_num")
	return mods, err
}

//获取所有用户信息
func UserAll() ([]User, error) {
	mods := make([]User, 0)
	err := DB.Select(&mods, "SELECT * FROM `goUser`") //select一个集合
	return mods, err
}

//获取一条信息
func UseOne() (User, error) {
	mods := User{}
	err := DB.Get(&mods, "SELECT * FROM goUser WHERE id = ?", 1)
	return mods, err
}

//
func Update() (sql.Result, error) {
	ret, err := DB.Exec("UPDATE goUser SET name = ? WHERE  id = ?", "test", 1)
	rows, _ := ret.RowsAffected()
	lastId, _ := ret.LastInsertId()
	fmt.Printf("Last inserted id is %v,row is %v", lastId, rows)
	return ret, err
}
