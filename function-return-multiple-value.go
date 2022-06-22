package main

import "fmt"

// function returning multiple values
func getFullName() (string, string, string) {
	return "Giri","Kusuma", "Made"
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
	
	name, _, _ := getFullName()
	fmt.Println(name)
}