package main

import (
	"fmt"
	"io"

	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println(" requests in golang")
	

	// Getreq()
	// postreq()
	performpostFOrm()
}


func Getreq() {
	
	const url = "http://localhost:3000/get"
	res, err := http.Get(url)
	if err != nil {
		panic(err)

	}
	defer res.Body.Close()
	fmt.Println("Statues code  is", res.StatusCode)
	fmt.Println("Lenth of the contemt", res.ContentLength)

	var responsesting strings.Builder
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))

	byteCount, _ := responsesting.Write(content)
	fmt.Println("BYte count is ", byteCount)
	fmt.Println(responsesting.String())
}


func postreq()  {
const url="http://localhost:3000/post"
requestBody:=strings.NewReader(`
{
	"name":"Archit",
	"Language":"Golang"

}
`)

	res,err:=http.Post(url,"application/json",requestBody)
	if err!=nil {
		panic(err)
	}
	defer res.Body.Close()

	contet,_:=io.ReadAll(res.Body)
	fmt.Println(string(contet))


}

func performpostFOrm()  {

	const url1="http://localhost:3000/postform"
	data:=url.Values{}
	data.Add("firstname","Archit")
	data.Add("Language","Golang")
	data.Add("email","archit.dev@gmail.com")
	res,err:=http.PostForm(url1,data)
	if err!=nil {
		panic(err)
		
	}

	defer res.Body.Close()

	content,_:=io.ReadAll(res.Body)
	fmt.Println(string(content))



}

