package main

import "fmt"

func main() {
	var a, b, d int
	var c string
	fmt.Println("input a and b:")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println("input you want correspondence of four operations:+,-,*,/")
	fmt.Scan(&c)
	switch c {
	case string('+'):
		d = a + b
		fmt.Println("result=", d)
	case string('-'):
		d = a - b
		fmt.Println("result=", d)
	case string('*'):
		d = a * b
		fmt.Println("result=", d)
	case string('/'):
		if b == 0 {
			fmt.Println("It's error")
		} else {
			d = a / b
			fmt.Println("result=", d)
		}

	}
}
