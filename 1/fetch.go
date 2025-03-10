// fetch 输出从URL获取的内容
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetchOne() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		count, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
		fmt.Printf("status code: %d, read size: %d\n", resp.StatusCode, count)
	}
}
