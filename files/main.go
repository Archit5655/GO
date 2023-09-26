package main

import (
	"fmt"
	"io"

	"os"
)

func main()  {
	fmt.Println("Files in golangg")
	content:="I am goin into a file "
	file , err :=os.Create("./createdfile.txt")
	if err!=nil {
		panic(err)


	}
	io.WriteString(file,content)
	readfile("./createdfile.txt")
	
}
func readfile(file string){
	 data,err:=os.ReadFile(file)
	 if err!=nil {
		fmt.Println(err)
		
	 }
	 fmt.Println(string(data))
	 




}