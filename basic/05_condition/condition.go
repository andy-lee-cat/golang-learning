package main

import (
	"fmt"
)

func main() {
	var val int
	fmt.Scanf("%d", &val)
	if val > 0 {
		fmt.Println("positive")
	} else if val < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero")
	}

	fmt.Scanf("%d", &val)
	switch val {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	default:
		fmt.Println("unknown")
	}

	var x interface{}
	switch i := x.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Printf("unknown: %T\n", i)
	}

	// equel to cpp switch, no break
	switch {
	case val > 0:
		fmt.Println("positive")
		fallthrough
	case val < 0:
		fmt.Println("negative")
		fallthrough
	default:
		fmt.Println("zero")
	}
}
