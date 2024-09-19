//  3. Packages are places that can be used to organize the program code that we create in Go-Lang.
//     Example import from helper folder
package main

import (
	"fmt"
	"linknau/helper"
)

func main() {
	result := helper.SayHello("Danu")
	fmt.Println(result)
}
