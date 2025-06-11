package main

import "fmt"

var (
	a = 10
	b =200
)

func add(x, y int) {
	fmt.Println("total of a and b: ", x+y)
}
func main() {
	fmt.Println("Hello, World! main")
	add(1, 2)
	add(a, 20)
	add(a, b)
}
func init() {
	fmt.Println("Hello, World!")
}

// * In Go (and many other programming languages), data segments refer to areas of a program's memory where global and static variables are stored
// * 1. What is the purpose of the init() function in Go? When is it called?
//* Ans ==>  The init() function is used to initialize variables or state before the main() function runs. It is called automatically by the Go runtime before main(), and you cannot call it explicitly.


// * Where is the variable a stored in memory, and why?
//* Ans ==> The variable a is stored in the data segment because it is declared at the package level (outside any function), making it a global variable.


// *  Explain the difference between global variables and local variables in terms of memory allocation.
//* Ans ==> Global variables are stored in the data segment and exist for the lifetime of the program. Local variables are stored on the stack and exist only during the execution of the function in which they are declared


//* What is a data segment, and what kind of variables are stored there in Go?
//* Ans ==> A data segment is a region of a program’s memory where global and static variables are stored. In Go, variables declared outside of functions (at the package level) are stored in the data segment

// * Can you have multiple init() functions in a Go program? If so, how are they executed?
// * Ans ==> Yes, you can have multiple init() functions in different files of the same package. They are executed in the order in which the files are compiled, and within a single file, in the order they appear

//  * What happens if you declare a variable inside the main() function instead of at the package level? Where is it stored?
//* Ans ==> If you declare a variable inside main(), it is a local variable and is stored on the stack. It only exists while main() is running.

//  * How does Go handle the initialization order of global variables and init() functions?
//* Ans ==> Go initializes global variables first, in the order they are declared, then calls the init() functions, and finally calls main().


// * Explain the difference between the stack, heap, and data segment in Go’s memory model.

//* Stack: Stores local variables and function call information; memory is managed automatically and is fast.
//* Heap: Stores dynamically allocated variables (e.g., via new or make); memory is managed by the garbage collector.
//* Data Segment: Stores global and static variables; exists for the lifetime of the program.
// * Code Segment: Stores executable instructions; exists for the lifetime of the program.

