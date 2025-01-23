package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions are here")

	greet()

	result := adder(3, 4)
	fmt.Println("result is this:", result)

	ProResult, message := proadder(2, 3, 4, 1)
	fmt.Println("Pro result is:", ProResult)
	fmt.Println(message)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proadder(val ...int) (int, string) {
	total := 0
	for _, sum := range val {
		total += sum
	}

	return total, "Here we end functions"
}

func greet() {
	fmt.Println("Namaste Banglore")
}
