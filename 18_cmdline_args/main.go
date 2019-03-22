package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("Usage: %s <filename.txt>\n", os.Args[0])
	}
	args := os.Args[1:]

	dataBytes, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(dataBytes))
}
