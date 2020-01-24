package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
func main() {
	dsn := "michael:cccbbb@tcp(192.168.11.31:30001)/spider_12306"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("faild,err:%v", err)
		return
	}
	defer db.Close()
	if err:=db.Ping();err!=nil{
		log.Fatalln(err)
	}
	row := db.QueryRow("select city from city_abbreviation where id = ?",4549 )
	var city string
	if err:=row.Scan(&city);err!=nil{
		log.Fatalln(err)
	}
	fmt.Println(city)

}
