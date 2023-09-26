package main
import (
	"fmt"
)


func main(){
	// fmt.Println("THis is the start of pointers")
	// var  value int
	// value=23;
	// var ptr *int	
	// fmt.Println("value of pointer is",ptr)
	mynum:=23
	//  &  is eqial to reference
	var ptr = &mynum
	fmt.Println(*ptr)
	fmt.Println(ptr)
//  *ptr is the actual  value and ptr is the address

*ptr *= 2
// 46
fmt.Println(*ptr)

	
}