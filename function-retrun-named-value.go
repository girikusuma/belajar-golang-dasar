package main

import "fmt"

// function return named value
func getFullName2() (firstName string, lastName string, age int) {
	firstName = "Giri"
	lastName = "Kusuma"
	age = 23

	return
}

func main() {
	a, b, c := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}