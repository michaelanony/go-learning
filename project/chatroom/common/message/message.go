package message

const(
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginRes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)
type Message struct {
	Type string `json:"type"`
	Data string	`json:"data"`
}

type LoginMes struct{
	UserId int	`json:"user_id"`
	UserPwd string	`json:"user_pwd"`
	UserName string	`json:"user_name"`
}

type LoginRes struct {
	Code int `json:"code"`//500 表示未注册，200 表示成功
	Error string	`json:"error"`
}

type RegisterMes struct {
	User User	`json:"userHandle"`
}

type RegisterResMes struct {
	Code int `json:"code"`//400 表示已有该用户，200 表示注册成功
	Error string	`json:"error"`

}