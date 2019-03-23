package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var links []string = []string{
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://golang.org",
	"http://amazon.com",
}

var wg sync.WaitGroup

func main() {
	fmt.Println()

	for _, link := range links {
		go checkLink(link)
		wg.Add(1)
	}
	wg.Wait()
}

func checkLink(link string) {

	resp, err := http.Get(link)
	content, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		log.Fatalln(err1)
	}
	fmt.Printf("%T\n", content)

	defer resp.Body.Close()
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!.")
	wg.Done()
}
