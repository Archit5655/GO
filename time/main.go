package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcoem to time ")


// time:=time.Now()

// fmt.Println(time)

// 2023-09-23 17:54:15.1348483 +0530 IST m=+0.002661901
// fmt.Println(time.Format("01-02-2006 Monday 15:04:05"))
// In this way we can format the date of the time yes 01-02-2006 is fixed
// output-
// 09-23-2023
// createtime:=time.Date(2020, time.July,23,23,23,23,23,time.UTC())

d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
year, month, day := d.Date()

fmt.Printf("year = %v\n", year)
fmt.Printf("month = %v\n", month)
fmt.Printf("day = %v\n", day)



}