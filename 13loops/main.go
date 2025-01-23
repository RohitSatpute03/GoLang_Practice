package main 

import (
	"fmt"
)

func main(){
	fmt.Println("LOOOOPSS")

	// days := []string{"SUnday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	// fmt.Println(days)

	// for d := 0; d < len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days{
	// 	fmt.Printf("At index %v day is %v\n", index, day)
	// }

	count := 0
	for count < 10{
		// if count == 4{
		// 	break
		// }

		// if count == 6{
		// 	count++
		// 	continue
		// }

		if count == 3{
			goto locd
		}

		fmt.Println("Count is: ",count)
		count++
	}

	locd:
		fmt.Println("THis is where we break")
}