package main

import "fmt"

func main() {
	fmt.Println("Strcuts")
	Archit := User{"Archit", "architagrgaQksjansd",true, 19}
	fmt.Println(Archit)
	fmt.Println(Archit.Email)
	fmt.Printf("THE details ofARchit are %+v\n",Archit)



}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
