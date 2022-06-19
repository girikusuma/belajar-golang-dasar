package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Cantiawan"
	names[1] = "Giri"
	names[2] = "Kusuma"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		17,
		10,
		98,
	}

	fmt.Println(values)

	fmt.Println(len(names))
}