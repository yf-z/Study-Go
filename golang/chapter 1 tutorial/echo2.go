// echo prints its command-line arguments
// version 2.0

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", "" // the type will be determined according to the initialization value
	// so in this case, s and sep will be strings
	for _, arg := range os.Args[1:] { // another type of loop using range
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}