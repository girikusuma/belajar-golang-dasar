package main

import "fmt"

// function variadic
func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total = total + value
	}

	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)

	slice := []int{20, 30, 40}
	total = sumAll(slice...)
	fmt.Println(total)
}