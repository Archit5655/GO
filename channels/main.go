package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("HEllo channels")
	mychannel:=make(chan int,2)
	wg:=&sync.WaitGroup{}

	//  <- this is used to pust value in golang go always ulta madarchod

	// mychannel <-5
	// fmt.Println(<-mychannel)
	wg.Add(2)

	go func(ch chan int,wg *sync.WaitGroup) {
		fmt.Println(<-mychannel)
		wg.Done()
	}(mychannel,wg)

	go func(ch chan int,wg *sync.WaitGroup) {
		mychannel<-5
		close(mychannel)

		
	
		wg.Done()
	}(mychannel,wg)


wg.Wait()
}
