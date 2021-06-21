// echo prints its command-line arguments
// version 3.0

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	//fmt.Println(os.Args[1:])
}