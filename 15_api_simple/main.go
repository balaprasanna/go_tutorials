package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const URL = "https://hub.helm.sh/api/chartsvc/v1/charts"

var cache ChartResponse

func root(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI, r.RemoteAddr, r.Header, r.Host, r.Method)
	fmt.Fprintln(w, "Hello world!")
}

func getchartsV1(w http.ResponseWriter, r *http.Request) {
	resp := Get(URL)
	fmt.Fprintln(w, resp)
}

func getchartsV2(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	resp := GetJSON(URL)
	jsonDataBytes, err := json.Marshal(resp)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(w, string(jsonDataBytes))
}

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", root)
	mux.HandleFunc("/charts", getchartsV2)
	mux.HandleFunc("/v1/charts", getchartsV1)
	mux.HandleFunc("/v2/charts", getchartsV2)

	fmt.Println(fmt.Sprintf("Starting app in :%s PORT", PORT))
	http.ListenAndServe(":"+PORT, mux)
}
