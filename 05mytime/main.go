package main 

import (
	"fmt"
	"time"
)

func main(){
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("Monday 01-02-2006 15:04:05 "))
}