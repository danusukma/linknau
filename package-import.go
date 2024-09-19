//  3. Pakage Management is Go’s package management is based on modules.
//     Each Go project is defined as a module, which can have dependencies on other modules.
//     The Go toolchain uses go.mod files to track dependencies.

//     How to Import:
//	   go mod init your-module-name
//     go get github.com/danusukma/linknau/helper

package main

import (
	"fmt"

	"github.com/danusukma/linknau/helper"
)

func main() {
	result := helper.SayHello("Danu")
	fmt.Println(result)
}
