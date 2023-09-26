package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case")


	// UNderstand this LATER
	
	rand.Seed(time.Now().UnixNano())

	dice := rand.Intn(6) + 1
	// fmt.Printf("Value of dice is %v ",dice)
	switch dice {
	case 1:
		fmt.Println("value is 1")
		break
	case 2:
		fmt.Println("valeyue is 2")

	case 3:
		fmt.Println("value 3")
		break
	case 4:

		fmt.Println("value 4")
		break
	default:
		fmt.Println("Laude iska tho  case bna")
		break
	}

}
