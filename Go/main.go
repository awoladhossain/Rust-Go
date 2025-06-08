package main

import "fmt"

// func main() {
// 	fmt.Println("hello awolad ")

//		add:=func(a int, b int) {
//			c := a + b
//			fmt.Println(c)
//		}
//		add(10,20)
//	}
func add(a, b int) {
	c := a + b
	fmt.Println(c)
}
func main() {
	fmt.Println("this is main function")
	add(10,20)
}


