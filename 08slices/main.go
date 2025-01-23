package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is for slices")

	var This_slice = []string{"Apple", "Orange", "Banana"}
	fmt.Printf("Type of slice is %T\n", This_slice)

	This_slice = append(This_slice, "Tomato", "Potato")
	fmt.Println("Complete slice is: ", This_slice)

	This_slice = append(This_slice[1:3])
	fmt.Println("we have sliced the slice: ",This_slice)

	highscores := make([]int, 4)
	highscores[0] = 982
	highscores[1] = 343
	highscores[2] = 1212
	highscores[3] = 232
	// highscores[4] = 0

	fmt.Println(highscores)
	highscores = append(highscores, 1212, 3233)
	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	//how to remove slices based on their index

	var courses = []string{"Java", "Python", "HTML", "goLang"}
	index := 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
