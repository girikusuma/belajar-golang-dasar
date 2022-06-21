package main

import "fmt"

func main(){
	var name string = "Giri"

	if name == "Giri" {
		fmt.Println("Hello Giri")
	} else if name == "Kusuma" {
		fmt.Println("Hello Kusuma")
	} else {
		fmt.Println("Hi")
	}

	// if short statement

	if length := len(name); length > 10 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}