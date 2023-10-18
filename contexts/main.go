package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// func doSomething(ctx context.Context) {
// 	fmt.Printf("Doing something and print the value of contrext %s\n ",ctx.Value("1"))
// 	fmt.Printf("Doing something and print the value of contrext %s\n ",ctx.Value("3"))
// 	fmt.Printf("Doing something and print the value of contrext %s\n ",ctx.Value("2"))
// }

// func doSomething(ctx context.Context) {
// 	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))

// 	anotherCtx := context.WithValue(ctx, "myKey", "anotherValue")
// 	doAnother(anotherCtx)

// 	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
// }

// func doAnother(ctx context.Context) {
// 	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value("myKey"))
// }

// func main() {
// 	// ctx := context.TODO()
// 	ctx := context.Background()

//		ctx=context.WithValue(ctx,"1","archit")
//		doAnother(ctx)
//		doSomething(ctx)
//	}
type Response struct {
	value int
	err   error
}

func fetchuserdata(ctx context.Context, userid int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respch := make(chan Response)

	go func() {

		value, err := fetchStuff()
		respch <- Response{
			value: value,
			err:   err,
		}

	}()
	for {
		select{
		case<-ctx.Done():
			return 0,fmt.Errorf("fetching datya took 2 long ")
		case resp:=<-respch:
			return resp.value,resp.err
		}
	}

}
func fetchStuff() (int, error) {
	time.Sleep(time.Millisecond * 170)
	return 666, nil

}

func main() {
	start := time.Now()
	ctx := context.Background()
	userid := 10
	val, err := fetchuserdata(ctx, userid)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("result: ", val)
	fmt.Println("took: ", time.Since(start))

}
