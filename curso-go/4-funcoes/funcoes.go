package main

import "fmt"

func somar(v1 int8, v2 int8) int8 {
	return v1 + v2
}

var f = func() string {
	return "função"
}

func main() {
	//primeiro código
	fmt.Println("primeiro código resultado", f)

	//segundo código
	fmt.Println("segundo código resultado", f())
}
