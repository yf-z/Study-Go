// Dup1 prints the text of each line that appears more than one in the standard input, preceded by its count
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // does not return false when the input is empty
		if (input.Text() == "") {
			break
		}
		counts[input.Text()]++
	}

	// Note: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}