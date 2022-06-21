package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{
		"name":    "Giri",
		"address": "Batur",
	}

	fmt.Println(person)

	character := map[string]string{
		"name": "Giri",
		"gender": "Male",
	}

	fmt.Println(character["name"])
	fmt.Println(len(character))

	book := make(map[string]string)

	book["title"] = "Belajar Go-Lang"
	book["author"] = "The Author"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
