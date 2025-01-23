package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to app")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter rating: ")

	input, _ := reader.ReadString('\n')
	fmt.Printf("Thank you for rating %v", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added +1 to your rating:", numRating+1)
	}
}
