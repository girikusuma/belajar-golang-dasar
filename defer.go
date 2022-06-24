package main

import "fmt"

// defer untuk memanggil function setelah function berjalanmeskipun terjadi error
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result =", result)
}

func main() {
	runApplication(0)
}