package main

import "fmt"

// Function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "***"
	} else {
		return name
	}
}

// Function as parameter with type declaration

type Filter func(string) string

func sayHelloWithFilter2(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter2(name string) string {
	if name == "Anjing" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Giri", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	sayHelloWithFilter2("Kusuma", spamFilter2)
	filter2 := spamFilter2
	sayHelloWithFilter2("Anjing", filter2)
}