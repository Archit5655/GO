package main
import (

	"fmt"
)


func main(){
	fmt.Println("Array in golang")
	var array [5] string
	array[0]="lauda"

	array[2]="java"
	array[3]="react"
	array[4]="c++"
	for i := 0; i < 5; i++ {
		fmt.Println(array[i])
		
	}
	fmt.Println(" ")
	fmt.Println(array)

	fmt.Println(len(array))
	var lang= [3] string {"Golang","Javascccript","JAva"}
	fmt.Println(lang)
	var num=[3] int {12,3,34}
	fmt.Println(num)
	
}



