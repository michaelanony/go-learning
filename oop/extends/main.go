package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score int
}

func (stu *Student) showInfo(){
	fmt.Println(stu.Age,stu.Name ,stu.Score)
}
type  Pupil struct {
	Student
}

func (p *Pupil) testing()  {
	fmt.Println("Pupil is testing")
}
type E struct {
	int //匿名
	n int
}

func main() {
	pupil := &Pupil{Student{}}
	pupil.Student.Name="michael"
	pupil.showInfo()
	pupil.testing()
	var e E
	e.int=30
	e.n=40
	fmt.Println(e)
}

