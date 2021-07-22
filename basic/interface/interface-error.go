package main

import (
	"errors"
	"fmt"
)

func Division(value int, division int) (int, error) {
	if division == 0 {
		return 0, errors.New("Division by zero")
	} else {
		result := value / division
		return result, nil
	}
}
func main() {
	result, err := Division(100, 0)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Hasil", result)
	}
}
