package main

import "fmt"

func Receiver(v interface{}) {
	switch v.(type) {
	case nil:
		fmt.Println("是个nil")
	case int:
		fmt.Println("是int")
	case string:
		fmt.Println("是string")
	case bool:
		fmt.Println("是bool")
	default:
		fmt.Println("是其他")
	}
}
func main() {
	var i string
	var a int = 32
	i = "你好吗"
	Receiver(i)
	Receiver(a)
}
