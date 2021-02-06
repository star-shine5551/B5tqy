package main

//字符串索引和其如何打印
import "fmt"

func main() {
	var a string = "hear"
	var b int
	b = len(a)
	c := byte(a[2])
	fmt.Printf("%d\n", b)
	fmt.Printf("%c\n", a[2])
	fmt.Println(c)
}
