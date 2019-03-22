package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter int
var PORT = ":8080"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter++
		fmt.Fprintln(w, "Count : "+strconv.Itoa(counter))
	})
	fmt.Println("App started in " + PORT)
	http.ListenAndServe(PORT, mux)
}
