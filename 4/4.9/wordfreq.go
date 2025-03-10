package main

import (
	"bufio"
	"os"
)

func wordfreq() map[string]int {
	counts := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		counts[in.Text()]++
	}
	return counts
}

func main() {
	counts := wordfreq()
	for word, count := range counts {
		println(word, count)
	}
}
