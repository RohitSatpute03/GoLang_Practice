package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Handling files in this ")

	newFile := "This is where im creating a file"
	file, err := os.Create("./locateFile.txt")
	// if err != nil{
	// 	panic(err)
	// }
	checkError(err)

	length, err := io.WriteString(file, newFile)
	checkError(err)

	fmt.Println("Length of the file is:", length)
	defer file.Close()

	readFile("./locateFile.txt")
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	checkError(err)
	fmt.Println("THis is the data in the file", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
