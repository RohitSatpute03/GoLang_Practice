package main

import (
	"fmt"
)

func main() {
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Tomato"

	fmt.Println("List of fruits: ", fruitList)
	fmt.Println("Length of the array is: ", len(fruitList))

	var vegies = [10]string{"a", "b", "c", "d"}
	fmt.Println("THese are the vegies: ", vegies)
	fmt.Println("Length of it is: ", len(vegies))

}
