package main

import "fmt"

func operation(a int, b int, op func(c int, d int)) {
    op(a, b)
}
func add(x,y int){
	fmt.Println("total of a and b: ", x+y)
}

func main() {
    operation(2, 5, add)
}
