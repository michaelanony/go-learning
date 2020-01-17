package main

import (
	"encoding/json"
	"go-learning/lyProject/model"
	"net/http"
)

func listAllLyPro(w http.ResponseWriter,r *http.Request)  {
	mods,_:= model.GetIntelPro()
	buf, _ := json.Marshal(mods)
	w.Header().Set("Content-Type", "application/json")
	_,err :=w.Write(buf)
	if err!=nil{
		panic(err)
	}
}
func main() {
	http.HandleFunc(`/lypro`,listAllLyPro)
	err:=http.ListenAndServe(":8082",nil)
	if err!=nil{
		panic(err)
	}
}