package main

import "fmt"

type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}
type Phone struct {

}
func (p Phone) Start()  {
	fmt.Println("Starting Phone...")
}
func (p Phone) Stop(){
	fmt.Println("Stopping Phone")
}
type Camera struct {

}
func (c Camera) Start()  {
	fmt.Println("Starting Camera...")
}
func (c Camera) Stop(){
	fmt.Println("Stopping Camera")
}

type Computer struct {

}
func (c Computer) Working (usb Usb){
	usb.Start()
	usb.Stop()
}
func main() {
	c := Computer{}
	phone := Phone{}
	camera := Camera{}
	c.Working(phone)
	c.Working(camera)
}