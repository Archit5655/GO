package main

import (
	"fmt"
)

func main() {
	fmt.Println("MEthods")
	Archit := User{"Archit", "architagrgaQksjansd", true, 19}
	fmt.Println(Archit)
	fmt.Println(Archit.Email)
	fmt.Printf("THE details ofARchit are %+v\n", Archit)
	Archit.GetStatus()

}

// ? Mehtods
func (u User) GetStatus() {
	fmt.Println("Is user active ", u.Status)


}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// age    int
}
