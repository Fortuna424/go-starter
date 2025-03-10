// dup output the duplicated lines that count > 1, counts at the begining
package main

import (
	"bufio"
	"fmt"
	"os"
)

func duplicateLines() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	// ps: ignore the error might in input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
