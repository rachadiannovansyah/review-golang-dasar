package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var names = []string{"Novan", "Rachadian"}
	printMessage("assalamualaikum", names)

	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(1, 100)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(1, 100)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(1, 100)
	fmt.Println("random number:", randomValue)
}

func randomWithRange(min, max int) int {
	var value = rand.Int() % (max - min)
	return value
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, "_")
	fmt.Println(message, nameString)
}
