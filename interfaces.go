//2. Interface is an abstract data type, it does not have a direct implementation.

package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Danu"}
	SayHello(person)

	animal := Animal{Name: "Dog"}
	SayHello(animal)
}
