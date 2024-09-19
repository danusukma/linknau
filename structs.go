//1. A struct is a collection of fields.

package main

import "fmt"

// Define struct
type Person struct {
	Nama string
	Age  int
}

func main() {
	fmt.Println("--- Soal No 1 --")

	//Create instant of struct
	MyPerson := Person{Nama: "Danu", Age: 42}

	//Read struct
	fmt.Println("My Name", MyPerson.Nama)
	fmt.Println("My Age", MyPerson.Age)
}
