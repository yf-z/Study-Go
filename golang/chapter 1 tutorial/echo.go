// echo prints its command-line arguments
// version 1.0
package main

import (
	"fmt"
	"os"
) // packages could be imported in the forms of lists or saparately like

// import "fmt"
// import "os"
// gofmt will sort all packages in alphabetical order

func main() {
	var s, sep string  // implicitly initialized as empty strings ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}