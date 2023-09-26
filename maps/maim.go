package main

import (
	"fmt"

)

func main(){
		fmt.Println("MAps in golang")
		lang:=make(map[int] string)
		lang[1]="React"
		lang[2]="JAVA"
		lang[3]="NODE"
		lang[4]="RUBY"
		lang[5]="RUST"
		fmt.Println(lang[1])
		fmt.Println(lang[3])
		delete(lang,2)
		fmt.Println(lang)

		for  _,value := range lang{
			fmt.Printf("FOr key v vlaue is %v\n",value)

		}



}