package main

import "fmt"

func main() {
	var a string = "hello"
	fmt.Println(a)

	var num1, num2 int = 1, 2
	fmt.Println(num1, num2)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var tval = true
	fmt.Printf("%v\n", tval)

	// intVal := 1 equal to var intVal int = 1
	var intVal int
	intVal = 1
	fmt.Printf("%v\n", intVal)
}
