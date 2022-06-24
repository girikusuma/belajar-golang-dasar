package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHalo(hasName HasName) {
	fmt.Println("Halo", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var giri = Person {
		Name: "Giri",
	}

	sayHalo(giri)

	var kucing = Animal {
		Name: "Kucing",
	}
	sayHalo(kucing)
}