package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://en.wikipedia.org/wiki/Go_(programming_language)"

func main() {
	fmt.Println("MOre about URLs ")

	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	// fmt.Println(result.Path)
	// fmt.Println(result.Host)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("THe type of qparams is %T\n", qparams)

	partsOfURL := &url.URL{
		Scheme: "https",
		Host:   "youtube.com",
	}

	constructed_url := partsOfURL.String()
	fmt.Println(constructed_url)
}
