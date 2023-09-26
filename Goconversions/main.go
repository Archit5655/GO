package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func  main()  {
	fmt.Println(" WElcome to  golang")
	fmt.Println("Plese ratre this chutiya language")
	reader:=bufio.NewReader(os.Stdin)
	rate,_:=reader.ReadString('\n')
	fmt.Printf("Thank you for rating us %v",rate )



	// rate is of type string
	rating,err:=strconv.ParseFloat(strings.TrimSpace(rate),64)
	if err !=nil {
		fmt.Println(err)		
	}	else {
		fmt.Println("Adeed 10 to your rate",rating+1)
	}
}