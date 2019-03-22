package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Data struct {
	Names []string `json:"names"`
	Age   []int    `json:"age"`
}

func main() {
	fmt.Println("JSON Tutorials")

	dataBytes, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Println(string(dataBytes))
	v := Data{}
	if err1 := json.Unmarshal(dataBytes, &v); err1 != nil {
		log.Fatalln(err1)
	}

	fmt.Println(v)

}
