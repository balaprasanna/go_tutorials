package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Any interface{}

// INTERNAL FUNCIONS

func get(url string) []byte {
	client := http.Client{}

	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		log.Fatalln(err1)
	}

	// Add HTTP Header to request
	httpHeaders := map[string]string{
		"user-agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
	}

	// Modify the above code to add new http headers
	for k, v := range httpHeaders {
		req.Header.Add(k, v)
	}

	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Fatalln(err2)
	}
	defer resp.Body.Close()

	data, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		log.Fatalln(err3)
	}
	return data
}

// Get : Http get util
func Get(url string) string {
	dataBytes := get(url)
	respStr := string(dataBytes)
	return respStr
}

// GetJSON : Http get util
func GetJSON(url string, data Any) Any {
	dataBytes := get(url)

	// var data Any
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
