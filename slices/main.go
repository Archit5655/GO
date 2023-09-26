package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in golnag")
	var lang = []string{"java", "Javascript", "GO", "RUST"}
	fmt.Printf("type of lang i s %T", lang)
	fmt.Println(" ")
	lang = append(lang, "c++", "c")
	fmt.Println(lang)
	lang = append(lang[1:3])
	fmt.Println(lang)

	var arr = []int{1, 2, 3, 4, 5}
	arr = append(arr[1:3])

	// make funciton

	score := make([]int, 5)
	score[0] = 234
	score[1] = 2134
	score[2] = 435
	score[3] = 4567
	score[4] = 654

	score = append(score, 2345, 5342456, 43, 22, 455)
	sort.Ints(score)
	fmt.Println(score)
	fmt.Println(sort.IntsAreSorted(score))

	//  removing the value of slice based on ideex

	var course = []string{"React", "Java", "GOLANG", "NODE", "frontend", "PHP"}
	fmt.Println(course)
	var i int = 3
	course = append(course[:i], course[i+1:]...)
	fmt.Println(course)
	
}
