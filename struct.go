package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
	Married bool
}

// struct function / method
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var giri Customer
	
	giri.Name = "Giri"
	giri.Address = "Batur"
	giri.Age = 23
	giri.Married = false
	
	// fmt.Println(giri)
	// fmt.Println(giri.Name)
	
	//memanggil struct method
	giri.sayHi("Kusuma")


	// struct literals
	// kusuma := Customer {
	// 	Name: "Kusuma",
	// 	Address: "Kintamani",
	// 	Age: 22,
	// 	Married: false,
	// }
	// fmt.Println(kusuma)

	// made := Customer{"Made", "Bangli", 32, true}
	// fmt.Println(made)
}
