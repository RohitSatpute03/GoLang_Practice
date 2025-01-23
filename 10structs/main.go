package main 
import (
	"fmt"
)

type User struct{
	Name string
	Email string 
	Status bool
	Age int
}

func main(){
	fmt.Println("This is about structs ")

	//no inheritance in golang
	Rohit := User{"Rohit", "rr@bb", true, 23}
	fmt.Println(Rohit)
	fmt.Printf("Rohit details are %+v: \n", Rohit)
	fmt.Printf("name is %v, email is %v, status is %v, age is %v\n", Rohit.Name, Rohit.Email, Rohit.Status, Rohit.Age)

}