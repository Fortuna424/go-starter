package main

import (
	"fmt"
	"math/rand"
	"time"

	// "the.go/tempconv"
	"the.go/popcount"
)

// func main() {
// 	for _, arg := range os.Args[1:] {
// 		t, err := strconv.ParseFloat(arg, 64)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
// 			os.Exit(1)
// 		}

// 		f := tempconv.Fahrenheit(t)
// 		c := tempconv.Celsius(t)
// 		k := tempconv.Kelvins(t)
// 		fmt.Printf("%s = %s, %s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c), k, tempconv.KToC(k))
// 	}
// }

func main() {
	start := time.Now()
	c := make(chan string)
	for i := 0; i < 2; i++ {
		go func(s time.Time, ch chan<- string) {
			num := rand.Uint64()
			count := popcount.PopCountCustom(num)
			ch <- fmt.Sprintf("%v popcount = %v, cost: %v", num, count, time.Since(s).Nanoseconds())
		}(start, c)
	}
	for i := 0; i < 2; i++ {
		fmt.Println(<-c)
	}
}
