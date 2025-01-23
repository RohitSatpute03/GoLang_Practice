package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	message := "Welcome to application"
	fmt.Println(message)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter feedback: ")

	//comma ok || err err

	input, _ := reader.ReadString('\n')
	fmt.Printf("Thank you for rating, %v", input)
	fmt.Printf("Data type of the input is %T\n", input)
}