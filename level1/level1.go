package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}
type dove interface {
	gugugu()
}
type repeater interface {
	repeat(string)
}
type NingMengJing interface {
	sour()
}
type ZhenXiangGuai interface {
	next(string)
}

func (a *person) next(word string) {
	fmt.Println(word, "这牢饭还行")
}
func (a *person) sour() {
	fmt.Println("你没鸽他，为什么要鸽我，我酸了！")
}
func (a *person) gugugu() {
	fmt.Println(a.name)
	fmt.Printf("我们是好兄弟，我怎么会鸽你呢！\n")
}
func (a *person) repeat(word string) {
	fmt.Println(word)
}
func main() {
	a := &person{
		name:   "DoveKing",
		age:    18,
		gender: "male",
	}
	a.gugugu()
	var r repeater
	r = a
	r.repeat("下次一定！")
	var c NingMengJing
	c = a
	c.sour()
	var d ZhenXiangGuai
	d = a
	d.next("我怎么可能喜欢萝莉呢！")
}
