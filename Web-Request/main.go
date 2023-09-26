package main

import (
	"fmt"
	"io"


	"net/http"
)

const url="https://lco.dev/"
func main(){
	fmt.Println("Webrequest")
res,err:=	http.Get(url)

if err!=nil {
	panic(err)
	
}
fmt.Printf("Type of respomse is %T",res)
// fmt.Println(res)


defer res.Body.Close() // ! close the connection in the end of the  program using defer

data,err:=io.ReadAll(res.Body)


if err!=nil {
	panic(err)
	
}
fmt.Println(string(data))


}