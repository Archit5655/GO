package main

import (
	"fmt"
)

func main() {
	fmt.Println("LOOOOOOOOPS")
	days := []string{"SUN", "MONDAY", "TUE", "WED", "THURS", "FRI", "SAT"}
	// fmt.Println(days)

	// ! LOOP 1

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// ! LOOP 2

	// for i :=range days{
	// 	fmt.Println(days[i])

	// }

	// ! Loop of type 3

	for index, day := range days {
		fmt.Printf("the index i s %v and the value is %v\n ", index+1, day)

	}

	// ! LOpp 4 THIS is a kind of while loop in golang
	value := 6
	for value < 10 {

		if value == 7 {
			// goto keyword jump directlt to the given statement
			goto lauda

		}
		fmt.Println("TU tho gaya laude")
		value++

	lauda:
		fmt.Println("agaya idaar")
		break

	}
}
