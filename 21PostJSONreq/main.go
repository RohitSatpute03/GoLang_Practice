package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("JSON post request")
	PerformPostjsonReq()
}

func PerformPostjsonReq() {
	const myurl = "http://localhost:8501/post"

	requestBody := strings.NewReader(`
		{
			"coursename":"golang",
			"faculty":"Rohit",
			"price":69000
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
