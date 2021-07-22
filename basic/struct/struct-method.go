package main

import "fmt"

type User struct {
	Id      int
	Name    string
	Age     int
	Married bool
}

// Struct function/method
func (user User) sayHello() {
	fmt.Println("Hello, my name is", user.Name)
	fmt.Println("my age is", user.Age)
}

func (user User) sayHi(name string) {
	fmt.Println("hello", name, "My name is", user.Name)
}

func main() {
	person := User{
		Name: "Novan",
		Age:  24,
	}

	person.sayHello()
	person.sayHi("Tika")
}
