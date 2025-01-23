package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("THis is switch cases")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(9) + 1
	fmt.Printf("Dice number is: %v\n", diceNum)

	switch diceNum{
	case 1:
		fmt.Println("1 YOu can open")
	case 2:
		fmt.Println("MOve 2 steps")
	case 3:
		fmt.Println("MOve 3 steps ")
	case 4:
		fmt.Println("MOve 4 steps")
		fallthrough
	case 5:
		fmt.Println("Move 5 steps")
	case 6:
		fmt.Println("it's a sixxx")
	default:
		fmt.Println("NOt valid !!!")
	}
}