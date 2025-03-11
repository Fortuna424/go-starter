package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Info struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

const xkcdUrl = "http://xkcd.com"

var wg sync.WaitGroup

func Download(ch chan int) {
	wg.Add(1)
	for i := range ch {
		url := fmt.Sprintf("%s/%d/info.0.json", xkcdUrl, i)
		resp, err := http.Get(url)
		if err != nil {
			log.Default().Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Default().Printf("HTTP request failed with status code %d", resp.StatusCode)
		}
		var info Info
		if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
			log.Default().Println(err)
		}
		// save to file
		file, err := os.Create(fmt.Sprintf("data/%d.json", i))
		if err != nil {
			log.Default().Println(err)
		}
		defer file.Close()
		enc := json.NewEncoder(file)
		enc.SetIndent("", "    ")
		if err := enc.Encode(info); err != nil {
			log.Default().Println(err)
		}
		log.Printf("Downloaded %d", i)
		time.Sleep(2 * time.Second)
	}
	wg.Done()
}

func main() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go Download(ch)
	}

	for i := 1; i <= 3060; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()

}
