package main 

import (
	"fmt"
)

func main(){
	fmt.Println("THis is if else condition")

	logincount := 23

	var result string
	if logincount < 10{
		result = "Regular user"
	} else if logincount > 10 {
		result = "Watch out"
	} else{
		result = "Exactly 10"
	}

	fmt.Println(result)


	if 8%2 == 0{
		fmt.Println("Even number")
	} else{
		fmt.Println("Odd number")
	}

	if num:=5; num < 10{
		fmt.Println("Less than 10")
	} else{
		fmt.Println("Greater than 10")
	}
}