package main

import "fmt"

func main() {
	name := "Kusuma"

	switch name {
	case "Giri":
		fmt.Println("Hello Giri")
	case "Kusuma":
		fmt.Println("Hello Kusuma")
	default:
		fmt.Println("Hi")
	}

	// switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	//switch tanpa value (kondisi)
	panjang := len(name)
	switch {
	case panjang > 10:
		fmt.Println("Nama terlalu panjang")
	case panjang > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}