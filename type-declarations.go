package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpGiri NoKTP = "5106041710980005"
	var marriedStatusGiri = false

	fmt.Println(noKtpGiri)
	fmt.Println(marriedStatusGiri)
}