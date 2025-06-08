package main

import (
	"fmt"
)

func main() {
	// int value
	// var a int =10
	// inferring type
	// var a=10
	// short variable declaration
	a := 10
	fmt.Println("Value of a:", a)

	// array
	// var b = [...]int{1, 2, 3, 4, 5}
	// var b[5]int = [5]int{1, 2, 3, 4, 5}
	// inferring type
	var b = [5]int{1, 2, 3, 4, 5}
	// short variable declaration
	// b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array b:", b)

	// slice means dynamic array
	var c = []int{1, 2, 3, 4, 5}
	// short variable declaration
	// c := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice c:", c)

	// map means key-value pair
	var d = map[string]int{"one": 1, "two": 2, "three": 3}
	// short variable declaration
	// d := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("Map d:", d)

	// struct means user-defined type
	type Person struct {
		Name string
		Age  int
	}
	var e = Person{Name: "Alice", Age: 30}
	// short variable declaration
	// e := Person{Name: "Alice", Age: 30}
	fmt.Println("Struct e:", e)

	var f = map[int]string{1: "awolad", 2: "anik"}
	fmt.Println("Map f:", f)

	type Car struct {
		Brand string
		Model string
		Year  int
	}
	var g = Car{Brand: "Toyota", Model: "Corolla", Year: 2020}
	// short variable declaration
	// g := Car{Brand: "Toyota", Model: "Corolla", Year: 2020}
	fmt.Println("Struct g:", g)

	// const means constant value
	const pi = 3.14
	fmt.Println("Value of pi:", pi)
	// const can also be used with type
	const euler float64 = 2.71828
	fmt.Println("Value of euler:", euler)
}
