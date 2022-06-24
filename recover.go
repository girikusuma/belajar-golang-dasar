package main

import "fmt"

func endApps() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
}

func runApps(error bool) {
	defer endApps()

	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApps(false)
	fmt.Println("Giri")
}