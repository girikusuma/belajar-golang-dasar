package main

import "fmt"

// function return value
func getHello(name string) string {
	return "Hello " + name
}

func main() {
	hello := getHello("Giri")
	fmt.Println(hello)
}