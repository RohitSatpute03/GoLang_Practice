package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://en.wikipedia.org/wiki/Go_(programming_language)"

func main(){
	fmt.Println("Web request handling in this")

	response , err := http.Get(url)
	checkError(err)

	fmt.Printf("Type of response is %T\n", response)

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	checkError(err)
	content := string(data)
	fmt.Println("Content inside the file is: ", content)
}

func checkError(err error){
	if err != nil{
		panic(err)
	}
}