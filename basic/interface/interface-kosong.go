package main

import "fmt"

func CobaInterfaceKosong(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return 2
	} else {
		return "Default gan"
	}
}

func main() {
	var data interface{} = CobaInterfaceKosong(3)

	fmt.Println(data)

	// same energy with
	data2 := CobaInterfaceKosong(3)
	fmt.Println(data2)
}
