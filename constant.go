package main

import "fmt"

func main(){
	const firstName string = "Giri"
	const lastName = "Kusuma"

	const (
		name = "Giri Kusuma"
		age = 23
	)

	fmt.Println(firstName)
	// fmt.Println(lastName)
	fmt.Println(name)
	fmt.Println(age)
}