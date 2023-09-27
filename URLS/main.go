package main

import (
	"fmt"
	"net/url"
)

const url1="https://lco.dev/learn?coursename=reactjs&paymentid=sniasjdfb"


func main(){
	fmt.Println("URLS IN GO")
	fmt.Println(url1)
result,_:=	url.Parse(url1)


fmt.Println(result.Fragment)
fmt.Println(result.Host)
fmt.Println(result.Path)
fmt.Println(result.Port())
fmt.Println(result.RawQuery)

qparams:=result.Query()
fmt.Printf("THE type of query parasms is %T\n",qparams)

fmt.Println(qparams)
fmt.Println(qparams["coursename"])

partsofurl:=&url.URL{
	Scheme:"https",
	Host:"lco.dev",
	Path:"/tutcss",
	RawPath:"user=hitesh",

}

anotherURL:=partsofurl.String()
fmt.Println(anotherURL)


}
