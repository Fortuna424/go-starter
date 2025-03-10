// echo implementation with go
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo() {
	fmt.Printf("time before run: %s\n", time.Now())
	for i, p := range os.Args {
		fmt.Printf("%d=%s\n", i, p)
	}
	fmt.Printf("time after run: %s\n", time.Now())
	fmt.Printf("time before run: %s\n", time.Now())
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Printf("time after run: %s\n", time.Now())
}
