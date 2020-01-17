package model

const (

)
type JsonStruct struct {}
type ReturnMes struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data string `json:"data"`
}
type MainConfig struct {
	User string
	Pwd  string
	Host string
	Port string
}
type lyIntellectualProperty struct {
	Id      int    `json:"id"`
	Iname   string `json:"iname"`
	Inameen string `json:"inameen"`
}

