package main

import "fmt"

func test(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Test"
	}
}

func main() {
	var result interface{} = test(1)
	fmt.Println(result)
}