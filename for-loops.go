package main

import "fmt"

func main() {
	// for
	counter := 1

	for counter <= 10 {
		fmt.Println("Counter ke -", counter)
		counter++
	}

	// for with statement
	for iteration := 1; iteration <= 10; iteration++ {
		fmt.Println("Iteration ke -", iteration)
	}

	sliceName := []string{"Giri", "Kusuma", "Made"}

	// for with index
	for i, value := range sliceName {
		fmt.Println("Index", i, "=", value)
	}

	// for without index
	for _, value := range sliceName {
		fmt.Println(value)
	}

	// map
	person := make(map[string]string)
	person["name"] = "Giri"
	person["title"] = "S.Kom"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}