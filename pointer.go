package main

import "fmt"

type Address struct {
	District, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// Not pointer
	// address1 := Address{"Bangli", "Bali", "Indonesia"}
	// address2 := address1
	
	// Pointer
	address1 := Address{"Bangli", "Bali", "Indonesia"} // var address1 Address = Address{"Bangli", "Bali", "Indonesia"}
	address2 := &address1 //var address2 *Address = &address1

	address2.District = "Denpasar"

	// address2 = &Address{"Jakarta", " DKIJakarta", "Indonesia"} // tidak merubah address1
	*address2 = Address{"Jakarta", " DKIJakarta", "Indonesia"} // merubah semua address yang mengacu ke Address yang sama

	fmt.Println(address1)
	fmt.Println(address2)
	
	// dengan function new
	var address3 *Address = new(Address)
	fmt.Println(address3)

	// Pointer di function
	var alamat = Address {
		District: "Badung",
		Province: "Bali",
		Country: "",
	}

	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

}