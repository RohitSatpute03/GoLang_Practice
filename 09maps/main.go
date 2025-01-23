package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is for maps (Keys and values)")

	This_map := make(map[string]string)
	This_map["py"] = "Python"
	This_map["rb"] = "Ruby"
	This_map["js"] = "JavaScript"

	fmt.Println("This is the map: ", This_map)
	fmt.Println("JS stands for: ", This_map["js"])

	delete(This_map, "rb")
	fmt.Println("Map after deletion: ", This_map)

	//loops in maps
	for key, value := range This_map {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
