package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is about pointers")

	// var ptr *int
	// fmt.Println(ptr)

	name := "Rohit"
	ptr := &name

	fmt.Println(ptr)
	fmt.Println(*ptr)

	var age int = 23
	ptr1 := &age
	*ptr1 = *ptr1 + 2
	fmt.Println(*ptr1)

}
