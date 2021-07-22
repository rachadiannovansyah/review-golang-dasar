package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Cimahi", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	// pointer function
	alamat := Address{
		City:     "Tangerang",
		Province: "Banten",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
