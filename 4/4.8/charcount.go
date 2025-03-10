package main

import "unicode"

const (
	LETTER = iota
	DIGIT
	OTHER
)

func charcount(s string) []int {
	counts := make([]int, 3)
	for _, r := range s {
		if unicode.IsLetter(r) {
			counts[LETTER]++
		} else if unicode.IsDigit(r) {
			counts[DIGIT]++
		} else {
			counts[OTHER]++
		}
	}
	return counts
}
