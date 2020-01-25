package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func indexView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func listView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("templates/list.html")
	w.Write(buf)
}
func userAll(w http.ResponseWriter, r *http.Request) {
	mods, _ := UserAll()
	buf, _ := json.Marshal(mods)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}
func userOne(w http.ResponseWriter, r *http.Request) {
	mods, _ := UseOne()
	buf, _ := json.Marshal(mods)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}
func update(w http.ResponseWriter, r *http.Request) {
	mods, _ := Update()
	buf, _ := json.Marshal(mods)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

func lyStaticNum(w http.ResponseWriter, r *http.Request) {
	mods, _ := lyStaticNumAll()
	buf, _ := json.Marshal(mods)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}
func main() {
	http.HandleFunc(`/`, listView)
	http.HandleFunc(`/userHandle`, userAll)
	http.HandleFunc(`/user1`, userOne)
	http.HandleFunc(`/userupdate`, update)
	http.HandleFunc(`/ly`, lyStaticNum)
	http.ListenAndServe(":8081", nil)
	fmt.Println("run")
}
