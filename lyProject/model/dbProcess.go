package model

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)

var DB *sqlx.DB

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}
func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
func init()  {
	var err error
	JsonParse:=NewJsonStruct()
	v := MainConfig{}
	JsonParse.Load("C:\\Users\\chenbo01\\Desktop\\code\\go-learning\\lyProject\\model\\config.json", &v)
	fmt.Println(v)
	DB, err = sqlx.Open(`mysql`,
		v.User+`:`+v.Pwd+`@tcp(`+v.Host+`:`+v.Port+`)/home?charset=utf8&parseTime=true`)
	if err != nil {
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	}
}
func GetIntelPro() ([]byte,error){
	var retMes ReturnMes
	mods := make([]lyIntellectualProperty,0)
	err := DB.Select(&mods,"SELECT * FROM `lywebback_intellectual_property`")
	if err!=nil{
		panic(err)
	}
	r,_ :=json.Marshal(&mods)
	retMes.Data=string(r)
	retMes.Code=200
	retMes.Message="ok"
	ret,_ :=json.Marshal(retMes)
	return ret,err
}