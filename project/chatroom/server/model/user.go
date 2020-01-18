package model

//先定义一个用户的结构体
type User struct {
	//为了序列和反序列化成功，必须保持用户信息的json字符串和结构体对应的tag名字一致
	UserId int	`json:"user_id"`
	UserPwd string `json:"user_pwd"`
	UserName string `json:"user_name"`
}
