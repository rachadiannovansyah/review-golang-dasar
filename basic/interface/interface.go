package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type User struct {
	Name string
}

func (user User) GetName() string {
	return user.Name
}

func main() {
	Person := User{
		Name: "Novan",
	}

	SayHello(Person)
}
