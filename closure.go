package main

import "fmt"

func main() {
	counter := 0
	name := "giri"

	increment := func() {
		fmt.Println("Increment")
		counter++

		name := "kusuma"
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}