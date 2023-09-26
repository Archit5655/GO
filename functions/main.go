package main

import (
	"fmt"
)

func main() {
	fmt.Println("Fuctions")
	sum(23, 23)
	p:=product(23,34)
	fmt.Println(p)
	sum:=addding(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(sum)
	sum2,mess:= addding2(2323,2,3,23,23)
	fmt.Println(sum2)
	fmt.Println(mess)


}

func sum(x int, y int) {
	fmt.Println(x + y)

}

func  product (x int ,y int ) int{
	return x*y
}
func addding (x ...int) int {
	sum:=0

	for _,value :=range x{
sum+=value

	}

	return sum
}


// ? FUnciton that returns two types of values
func addding2 (x ...int) (int,string) {
	

	sum:=0

	for _,value :=range x{
sum+=value

	}

	return sum,"TERI maka bsdsa"
}