package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var links []string = []string{
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://golang.org",
	"http://amazon.com",
}

func main() {
	fmt.Println()

	c := make(chan string)
	contentChan := make(chan string)
	contents := []string{}
	status := []string{}

	for _, link := range links {
		go checkLink(link, c, contentChan)
	}

	// for data := range c {
	// 	fmt.Println(data)
	// }
	// for index, _ := range links {
	// 	fmt.Println(index, <-c)
	// }

	for {
		select {
		case content := <-contentChan:
			contents = append(contents, content)
		case stat := <-c:
			status = append(status, stat)
		}
		if len(contents) == 5 {
			break
		}
	}
	for _, result := range status {
		fmt.Println(result)
	}
}

func checkLink(link string, c chan string, c1 chan string) {

	resp, err := http.Get(link)
	content, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		log.Fatalln(err1)
		c1 <- ""
	} else {
		c1 <- string(content)
	}

	defer resp.Body.Close()
	if err != nil {
		c <- fmt.Sprintf("%s - %s ", link, "might be down!")
		return
	}
	c <- fmt.Sprintf("%s - %s ", link, "is up!")
}
