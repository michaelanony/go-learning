package main

import ("fmt"
	"math/rand"
	"sort")

//1：声明Hero结构体
type Hero struct{
	Name string
	Age int
}
//2：声明一个Hero结构体切片类型
type HeroSlice []Hero
//3：实现interface接口
func (hs HeroSlice) Len() int{
	return len(hs)
}
//4:Less方法就是决定你是用什么标准进行排序
//4.1按Hero对年龄从小到大排序
func (hs HeroSlice) Less(i,j int) bool{
	return hs[i].Age > hs[j].Age
}
func (hs HeroSlice) Swap(i,j int) {
	 hs[i].Age,hs[j].Age = hs[j].Age,hs[i].Age
}
func main() {

	//定义一个数组/切片
	var intSlice=[]int{0,-1,10,7,90}
	//要求对intslice切片进行排序
	//1：冒泡排序
	//2：sort方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//对结构体进行排序
	var heroes HeroSlice
	for i :=0;i<10;i++{
		hero:=Hero{
			Name:fmt.Sprintf("Hero~%d",rand.Intn(100)),
			Age:rand.Intn(100),
		}
		heroes = append(heroes,hero)
	}
	for _,v:=range heroes{
		fmt.Println(v)
	}
	sort.Sort(heroes)
	fmt.Println(heroes)
	for _,v:=range heroes{
		fmt.Println(v)
	}
}