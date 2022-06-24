package main

import "fmt"

// No function recursive (using loop)
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// Function recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return value
	} else {
		return value * factorialRecursive(value - 1)
	}
}

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)

	recursive := factorialRecursive(10)
	fmt.Println(recursive)
}