package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var music int;
	// var ismusic bool;
	// ismusic=true;
	// // var lauda int;
	// music=23223233242;
	// fmt.Println(music)
	// fmt.Println(ismusic)
	// var lauda=23234;
	// fmt.Println(lauda)

	// fmt.Printf("the type of lauda is %T",lauda)
	welcome := "welcome to golang"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number : ")
	// Comma ok // err oK  syntax
	input, _ := reader.ReadString('\n') //it will takwe input and read the value of the string until it reacive \n or enter

	// The underscore (_) on the right side of the assignment is used to ignore the potential error returned by ReadString. In a production program, you would typically handle errors more gracefully.

	fmt.Println("Enter your First name")
	name, _ := reader.ReadString('\n')

	fmt.Println("your input is ", input)

	fmt.Printf("YOur name is %v", name)
	fmt.Printf("Type of name i is %T", name)

}
