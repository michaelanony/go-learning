package main

var captchaKey,phoneNum string

func main()  {
	if err:=InitPool();err!=nil{
		panic(err)
	}
	view()
}

