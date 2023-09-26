package main


import(
	"fmt"
)

func main(){
	defer fmt.Println("HEllo wrold")
	defer fmt.Println("one")
	defer fmt.Println("two")

	fmt.Println("DEFER")
	mydefer()


}
func mydefer(){
	for i := 0; i < 5; i++ {
	defer	fmt.Println(i)

		
	}
}