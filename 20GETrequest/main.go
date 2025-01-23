package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("WEB VERB-VIDEO")

	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "https://localhost:3000/get"

	response, err := http.Get(myurl)

	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code is:", response.StatusCode)
	fmt.Println("Length of content:", response.ContentLength)

	
}