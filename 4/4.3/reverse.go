package main

func reverse(slice *[]int) *[]int {
	for i, j := 0, len(*slice)-1; i < j; i, j = i+1, j-1 {
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
	return slice
}

func rotate(slice *[]int, n int) *[]int {
	*slice = append((*slice)[n:], (*slice)[:n]...)
	return slice
}

func dedup(slice *[]string) *[]string {
	i := 0
	for _, s := range *slice {
		if (*slice)[i] != s {
			i++
			(*slice)[i] = s
		}
	}
	*slice = (*slice)[:i+1]
	return slice
}

func reverseRune(slice *[]byte) *[]byte {
	for i, j := 0, len(*slice)-1; i < j; i, j = i+1, j-1 {
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
	return slice
}
