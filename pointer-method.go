package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	giri := Man{"Giri"}
	giri.Married()

	fmt.Println(giri.Name)
}