package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)


var signals=[]string{"test"}

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	// go greeter("laude")
	greeter("nikal")

	//  Anonymous function ---- this is the type of function in golang that is Anonymous functions are useful when you want to create an inline function or when a function is only going to be used once in your code12. They can also form closures, which means they can access variables from the outer function2.
	// func ()  {
	// 	fmt.Println("Hello world")
	// }()
	websiteslist :=[] string{
		"http://lco.dev",
		"http://go.dev",
		"http://google.com",
		"http://github.com",
	}

	
	

	for _,web:=range websiteslist{
		go getStatusCode(web)
		 wg.Add(1)

	}
	// this will not allow to finish the main method beforre the goroutines come back
	wg.Wait()
	fmt.Println(signals)

}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Millisecond)
		fmt.Println(s)

	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("error ")

	}
	mut.Lock()
	signals=append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status  code for %s\n", res.StatusCode, endpoint)

}
