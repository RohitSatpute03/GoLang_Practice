package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("WEB VERB-VIDEO")

	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8501/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code is:", response.StatusCode)
	fmt.Println("Length of content:", response.ContentLength)

	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	bytecode, _ := responseString.Write(content)

	fmt.Println("number of bytes:", bytecode)
	fmt.Println("Content inside it:", responseString.String())

	if err != nil {
		panic(err)
	}
	// fmt.Println(content)
	// fmt.Println(string(content))

}
