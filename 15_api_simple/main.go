package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const URL = "https://hub.helm.sh/api/chartsvc/v1/charts"

// Any ... generic interface hack
// type Any interface{}

func root(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI, r.RemoteAddr, r.Header, r.Host, r.Method)
	fmt.Fprintln(w, "Hello world!")
}

func getchartsV1(w http.ResponseWriter, r *http.Request) {
	resp := Get(URL)
	fmt.Fprintln(w, resp)
}

func getchartsV2(w http.ResponseWriter, r *http.Request) {

	// var data Temp
	resp := GetJSON(URL)
	fmt.Fprintln(w, resp)
}

func getcharts(w http.ResponseWriter, r *http.Request) {

	responseType := r.FormValue("type")

	client := http.Client{}

	req, reqErr := http.NewRequest("GET", URL, nil)
	if reqErr != nil {
		log.Fatalln(reqErr)
	}

	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Fatalln(reqErr)
	}

	defer resp.Body.Close()

	if responseType == "json" {
		decoder := json.NewDecoder(resp.Body)
		var data Any
		errDecode := decoder.Decode(&data)

		if errDecode != nil {
			log.Fatalln(errDecode)
		}

		fmt.Fprintln(w, data)
	} else {
		bodyBytes, errBody := ioutil.ReadAll(resp.Body)
		if errBody != nil {
			log.Fatalln(errBody)
		}
		fmt.Fprintln(w, string(bodyBytes))
	}

}

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	http.HandleFunc("/", root)
	http.HandleFunc("/charts", getcharts)
	http.HandleFunc("/v1/charts", getchartsV1)
	http.HandleFunc("/v2/charts", getchartsV2)

	fmt.Println(fmt.Sprintf("Starting app in :%s PORT", PORT))
	http.ListenAndServe(":"+PORT, nil)
}
